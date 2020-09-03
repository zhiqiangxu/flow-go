package chunks

import (
	"math/rand"
	"testing"

	"github.com/dapperlabs/flow-go/crypto/hash"
	"github.com/dapperlabs/flow-go/crypto/random"
	"github.com/dapperlabs/flow-go/model/encoding"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	chmodels "github.com/dapperlabs/flow-go/model/chunks"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/network/gossip/libp2p/test"
	protocolMock "github.com/dapperlabs/flow-go/state/protocol/mock"
	"github.com/dapperlabs/flow-go/utils/unittest"
)

// PublicAssignmentTestSuite contains tests against methods of the PublicAssignment scheme
type PublicAssignmentTestSuite struct {
	suite.Suite

	rngByBlockID func(flow.Identifier) (random.Rand, error)
}

func (a *PublicAssignmentTestSuite) SetupTest() (*flow.Header, *protocolMock.Snapshot, *protocolMock.State) {
	participants, _, _ := unittest.CreateNParticipantsWithMyRole(
		flow.RoleVerification,
		flow.RoleVerification,
	)

	// setup protocol state
	block, snapshot, state := unittest.FinalizedProtocolStateWithParticipants(participants)
	head := block.Header

	a.rngByBlockID = CreateRNGByBlockIDClosure(state)

	return head, snapshot, state
}

// TestAssignment invokes all the tests in this test suite
func TestAssignment(t *testing.T) {
	suite.Run(t, new(PublicAssignmentTestSuite))
}

// TestByNodeID evaluates the correctness of ByNodeID method of PublicAssignment
func (a *PublicAssignmentTestSuite) TestByNodeID() {
	size := 5
	// creates ids and twice chunks of the ids
	ids := test.CreateIDs(size)
	chunks := a.CreateChunks(2*size, a.T())
	assignment := chmodels.NewAssignment()

	// assigns two chunks to each verifier node
	// j keeps track of chunks
	j := 0
	for i := 0; i < size; i++ {
		c, ok := chunks.ByIndex(uint64(j))
		require.True(a.T(), ok, "chunk out of range requested")
		assignment.Add(c, append(assignment.Verifiers(c), ids[i].NodeID))
		j++
		c, ok = chunks.ByIndex(uint64(j))
		require.True(a.T(), ok, "chunk out of range requested")
		assignment.Add(c, append(assignment.Verifiers(c), ids[i].NodeID))
	}

	// evaluating the chunk assignment
	// each verifier should have two certain chunks based on the assignment
	// j keeps track of chunks
	j = 0
	for i := 0; i < size; i++ {
		assignedChunks := assignment.ByNodeID(ids[i].NodeID)
		require.Len(a.T(), assignedChunks, 2)
		c, ok := chunks.ByIndex(uint64(j))
		require.True(a.T(), ok, "chunk out of range requested")
		require.Contains(a.T(), assignedChunks, c.Index)
		j++
		c, ok = chunks.ByIndex(uint64(j))
		require.True(a.T(), ok, "chunk out of range requested")
		require.Contains(a.T(), assignedChunks, c.Index)
	}

}

// TestAssignDuplicate tests assign Add duplicate verifiers
func (a *PublicAssignmentTestSuite) TestAssignDuplicate() {
	size := 5
	// creates ids and twice chunks of the ids
	var ids flow.IdentityList = test.CreateIDs(size)
	chunks := a.CreateChunks(2, a.T())
	assignment := chmodels.NewAssignment()

	// assigns first chunk to non-duplicate list of verifiers
	c, ok := chunks.ByIndex(uint64(0))
	require.True(a.T(), ok, "chunk out of range requested")
	assignment.Add(c, ids.NodeIDs())
	require.Len(a.T(), assignment.Verifiers(c), size)

	// duplicates first verifier, hence size increases by 1
	ids = append(ids, ids[0])
	require.Len(a.T(), ids, size+1)
	// assigns second chunk to a duplicate list of verifiers
	c, ok = chunks.ByIndex(uint64(1))
	require.True(a.T(), ok, "chunk out of range requested")
	assignment.Add(c, ids.NodeIDs())
	// should be size not size + 1
	require.Len(a.T(), assignment.Verifiers(c), size)
}

// TestPermuteEntirely tests permuting an entire IdentityList against
// randomness and deterministicity
func (a *PublicAssignmentTestSuite) TestPermuteEntirely() {
	head, snapshot, _ := a.SetupTest()

	// create seed
	result := a.CreateResult(head, 4, a.T())
	seed := a.HashResult(result, a.T())

	snapshot.On("Seed", mock.Anything, mock.Anything, mock.Anything).Return(seed, nil)

	// creates random ids
	count := 10
	var idList flow.IdentityList = test.CreateIDs(count)
	var ids flow.IdentifierList = idList.NodeIDs()
	original := make(flow.IdentifierList, count)
	copy(original, ids)

	// Randomness:
	rng1, err := a.rngByBlockID(head.ID())
	require.NoError(a.T(), err)
	err = rng1.Shuffle(len(ids), ids.Swap)
	require.NoError(a.T(), err)

	// permutation should not change length of the list
	require.Len(a.T(), ids, count)

	// list should be permuted
	require.NotEqual(a.T(), ids, original)

	// Deterministiciy:
	// shuffling same list with the same seed should generate the same permutation
	rng2, err := a.rngByBlockID(head.ID())
	require.NoError(a.T(), err)
	// permutes original list with the same seed
	err = rng2.Shuffle(len(original), original.Swap)
	require.NoError(a.T(), err)
	require.Equal(a.T(), ids, original)
}

// TestPermuteSublist tests permuting an a sublist of an
// IdentityList against randomness and deterministicity
func (a *PublicAssignmentTestSuite) TestPermuteSublist() {
	head, snapshot, _ := a.SetupTest()

	// create seed
	result := a.CreateResult(head, 4, a.T())
	seed := a.HashResult(result, a.T())

	snapshot.On("Seed", mock.Anything, mock.Anything, mock.Anything).Return(seed, nil)

	// creates random ids
	count := 10
	subset := 4

	var idList flow.IdentityList = test.CreateIDs(count)
	var ids flow.IdentifierList = idList.NodeIDs()
	original := make([]flow.Identifier, count)
	copy(original, ids)

	// create result seed with 4 chunks
	result = a.CreateResult(head, 4, a.T())
	snapshot.On("Seed", mock.Anything, mock.Anything, mock.Anything).Return(seed, nil)

	// Randomness:
	rng1, err := a.rngByBlockID(head.ID())
	require.NoError(a.T(), err)
	err = rng1.Samples(len(ids), subset, ids.Swap)
	require.NoError(a.T(), err)

	// permutation should not change length of the list
	require.Len(a.T(), ids, count)

	// the initial subset of the list that is permuted should
	// be different than the original
	require.NotEqual(a.T(), ids[:subset], original[:subset])
}

// TestDeterministicy evaluates deterministic behavior of chunk assignment when
// chunks, random generator, and nodes are the same
func (a *PublicAssignmentTestSuite) TestDeterministicy() {
	head, snapshot, _ := a.SetupTest()

	c := 10    // keeps number of chunks
	n := 10    // keeps number of verifier nodes
	alpha := 1 // each chunk requires alpha verifiers

	// create seed
	result := a.CreateResult(head, c, a.T())
	seed := a.HashResult(result, a.T())
	snapshot.On("Seed", mock.Anything, mock.Anything, mock.Anything).Return(seed, nil)

	// creates two set of the same nodes
	nodes1 := test.CreateIDs(n)
	nodes2 := make([]*flow.Identity, n)
	require.Equal(a.T(), copy(nodes2, nodes1), n)

	// chunk assignment of the first set
	a1, err := NewPublicAssignment(alpha, a.rngByBlockID)
	require.NoError(a.T(), err)
	p1, err := a1.Assign(nodes1, result.Chunks, head.ID())
	require.NoError(a.T(), err)

	// chunk assignment of the second set
	a2, err := NewPublicAssignment(alpha, a.rngByBlockID)
	require.NoError(a.T(), err)
	p2, err := a2.Assign(nodes1, result.Chunks, head.ID())
	require.NoError(a.T(), err)

	// list of nodes should get shuffled after public assignment
	// but it should contain same elements
	require.Equal(a.T(), p1, p2)
}

// TestChunkAssignmentOneToOne evaluates chunk assignment against
// several single chunk to single node assignment
func (a *PublicAssignmentTestSuite) TestChunkAssignmentOneToOne() {
	// covers an edge case assigning 1 chunk to a single verifier node
	a.ChunkAssignmentScenario(1, 1, 1)

	// assigning 10 chunks to one node
	a.ChunkAssignmentScenario(10, 1, 1)
	// assigning 10 chunks to 2 nodes
	// each chunk to one verifier
	a.ChunkAssignmentScenario(10, 2, 1)
	// each chunk to 2 verifiers
	a.ChunkAssignmentScenario(10, 2, 2)

	// assigning 10 chunks to 10 nodes
	// each chunk to one verifier
	a.ChunkAssignmentScenario(10, 10, 1)
	// each chunk to 6 verifiers
	a.ChunkAssignmentScenario(10, 10, 6)
	// each chunk to 9 verifiers
	a.ChunkAssignmentScenario(10, 10, 9)
}

// TestChunkAssignmentOneToMay evaluates chunk assignment
func (a *PublicAssignmentTestSuite) TestChunkAssignmentOneToMany() {
	//  against assigning 52 chunks to 7 nodes
	//  each chunk to 5 verifiers
	a.ChunkAssignmentScenario(52, 7, 5)
	//  against assigning 49 chunks to 9 nodes
	//  each chunk to 8 verifiers
	a.ChunkAssignmentScenario(52, 9, 8)
}

// ChunkAssignmentScenario is a test helper that creates chunkNum chunks, verNum verifiers
// and then assign each chunk to alpha randomly chosen verifiers
// it also evaluates that each chuck is assigned to alpha many unique verifier nodes
func (a *PublicAssignmentTestSuite) ChunkAssignmentScenario(chunkNum, verNum, alpha int) {
	head, snapshot, _ := a.SetupTest()

	result := a.CreateResult(head, chunkNum, a.T())
	seed := a.HashResult(result, a.T())
	snapshot.On("Seed", mock.Anything, mock.Anything, mock.Anything).Return(seed, nil)

	// creates nodes and keeps a copy of them
	nodes := test.CreateIDs(verNum)
	original := make([]*flow.Identity, verNum)
	require.Equal(a.T(), copy(original, nodes), verNum)

	a1, err := NewPublicAssignment(alpha, a.rngByBlockID)
	require.NoError(a.T(), err)
	p1, err := a1.Assign(nodes, result.Chunks, head.ID())
	require.NoError(a.T(), err)

	// list of nodes should get shuffled after public assignment
	require.ElementsMatch(a.T(), nodes, original)

	for _, chunk := range result.Chunks {
		// each chunk should be assigned to alpha verifiers
		require.Equal(a.T(), p1.Verifiers(chunk).Len(), alpha)
	}
}

func (a *PublicAssignmentTestSuite) TestCacheAssignment() {
	head, snapshot, _ := a.SetupTest()

	result := a.CreateResult(head, 20, a.T())
	seed := a.HashResult(result, a.T())
	snapshot.On("Seed", mock.Anything, mock.Anything, mock.Anything).Return(seed, nil)

	// creates nodes and keeps a copy of them
	nodes := test.CreateIDs(5)
	assigner, err := NewPublicAssignment(3, a.rngByBlockID)
	require.NoError(a.T(), err)

	// initially cache should be empty
	require.Equal(a.T(), assigner.Size(), uint(0))

	// new assignment should be cached
	// random generators are stateful and we need to
	// generate a new one if we want to have the same
	// state
	_, err = assigner.Assign(nodes, result.Chunks, head.ID())
	require.NoError(a.T(), err)
	require.Equal(a.T(), assigner.Size(), uint(1))

	// repetitive assignment should not be cached
	_, err = assigner.Assign(nodes, result.Chunks, head.ID())
	require.NoError(a.T(), err)
	require.Equal(a.T(), assigner.Size(), uint(1))

	// creates a new set of nodes, hence assigner should cache new assignment
	newNodes := test.CreateIDs(6)
	_, err = assigner.Assign(newNodes, result.Chunks, head.ID())
	require.NoError(a.T(), err)
	require.Equal(a.T(), assigner.Size(), uint(2))

	// performs the assignment using a different seed
	// should results in a different new assignment
	// which should be cached
	otherResult := a.CreateResult(head, 20, a.T())

	_, err = assigner.Assign(newNodes, otherResult.Chunks, head.ID())
	require.NoError(a.T(), err)
	require.Equal(a.T(), assigner.Size(), uint(3))
}

// CreateChunk creates and returns num chunks. It only fills the Index part of
// chunks to make them distinct from each other.
func (a *PublicAssignmentTestSuite) CreateChunks(num int, t *testing.T) flow.ChunkList {
	list := flow.ChunkList{}
	for i := 0; i < num; i++ {
		// creates random state for each chunk
		// to provide random ordering after sorting
		state := make([]byte, 64)
		_, err := rand.Read(state)
		require.NoError(t, err)

		// creates chunk
		c := &flow.Chunk{
			Index: uint64(i),
			ChunkBody: flow.ChunkBody{
				StartState: state,
			},
		}
		list.Insert(c)
	}
	require.Equal(a.T(), num, list.Len())
	return list
}

func (a *PublicAssignmentTestSuite) CreateResult(head *flow.Header, num int, t *testing.T) *flow.ExecutionResult {
	list := a.CreateChunks(5, a.T())
	result := &flow.ExecutionResult{
		ExecutionResultBody: flow.ExecutionResultBody{
			BlockID: head.ID(),
			Chunks:  list,
		},
	}

	return result
}

func (a *PublicAssignmentTestSuite) HashResult(res *flow.ExecutionResult, t *testing.T) []byte {
	h := hash.NewSHA3_384()

	// encodes result approval body to byte slice
	b, err := encoding.DefaultEncoder.Encode(res.ExecutionResultBody)
	require.NoError(a.T(), err)

	// takes hash of result approval body
	hash := h.ComputeHash(b)

	return hash
}
