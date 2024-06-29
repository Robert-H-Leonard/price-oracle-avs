// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractPriceAggregatorTaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IPriceAggregatorTaskManagerAggregatedPrice is an auto generated low-level Go binding around an user-defined struct.
type IPriceAggregatorTaskManagerAggregatedPrice struct {
	Price             uint32
	Decimals          uint8
	LastBlockUpdated  uint32
	LastUpdatedTaskId uint32
}

// IPriceAggregatorTaskManagerPriceUpdateTask is an auto generated low-level Go binding around an user-defined struct.
type IPriceAggregatorTaskManagerPriceUpdateTask struct {
	TaskCreatedBlock          uint32
	FeedName                  string
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
	MinNumOfOracleNetworks    uint8
}

// IPriceAggregatorTaskManagerPriceUpdateTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IPriceAggregatorTaskManagerPriceUpdateTaskResponse struct {
	Price    uint32
	Decimals uint32
	TaskId   uint32
	Source   string
}

// IPriceAggregatorTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IPriceAggregatorTaskManagerTask struct {
	NumberToBeSquared         *big.Int
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IPriceAggregatorTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IPriceAggregatorTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	NumberSquared      *big.Int
}

// IPriceAggregatorTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IPriceAggregatorTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractPriceAggregatorTaskManagerMetaData contains all meta data concerning the ContractPriceAggregatorTaskManager contract.
var ContractPriceAggregatorTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_QUORUM_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_REQUIRED_NUM_OF_SOURCES\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"numberToBeSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fetchLatestAggregatedPrice\",\"inputs\":[{\"name\":\"feedName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPriceAggregatorTaskManager.AggregatedPrice\",\"components\":[{\"name\":\"price\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"lastBlockUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"lastUpdatedTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isValidOperator\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIPriceAggregatorTaskManager.Task\",\"components\":[{\"name\":\"numberToBeSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIPriceAggregatorTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numberSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIPriceAggregatorTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestPriceFeedUpdate\",\"inputs\":[{\"name\":\"feedName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"minNumOfOracleNetworks\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIPriceAggregatorTaskManager.PriceUpdateTask\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feedName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"minNumOfOracleNetworks\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"taskResponses\",\"type\":\"tuple[]\",\"internalType\":\"structIPriceAggregatorTaskManager.PriceUpdateTaskResponse[]\",\"components\":[{\"name\":\"price\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"decimals\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"source\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"nonSignerStakesAndSignatures\",\"type\":\"tuple[]\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature[]\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setQuorumThreshold\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRequiredNumberOfSources\",\"inputs\":[{\"name\":\"requiredNumOfSources\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPriceAggregatorTaskManager.Task\",\"components\":[{\"name\":\"numberToBeSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceUpdateRequested\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPriceAggregatorTaskManager.PriceUpdateTask\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feedName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"minNumOfOracleNetworks\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPriceAggregatorTaskManager.PriceUpdateTaskResponse\",\"components\":[{\"name\":\"price\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"decimals\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"source\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPriceAggregatorTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x610120604052604360c955600260ca553480156200001c57600080fd5b506040516200624b3803806200624b8339810160408190526200003f9162000201565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000099573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000bf919062000248565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000117573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013d919062000248565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000197573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001bd919062000248565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff1661010052506200026f565b6001600160a01b0381168114620001fe57600080fd5b50565b600080604083850312156200021557600080fd5b82516200022281620001e8565b602084015190925063ffffffff811681146200023d57600080fd5b809150509250929050565b6000602082840312156200025b57600080fd5b81516200026881620001e8565b9392505050565b60805160a05160c05160e05161010051615f2e6200031d600039600081816102a90152818161062501526130ce0152600081816105d201526122e50152600081816103dc01526124c701526000818161041b0152818161269d015261285f01526000818161046901528181610e9601528181611a3f01528181611fb0015281816121480152818161238201528181612b4001528181612e950152818161365c01526136f60152615f2e6000f3fe608060405234801561001057600080fd5b50600436106102485760003560e01c806372d18e8d1161013b578063cefdc1d4116100b8578063f2fde38b1161007c578063f2fde38b14610610578063f5c9899d14610623578063f63c5bab14610649578063f8c8765e14610651578063fabc1cbc1461066457600080fd5b8063cefdc1d414610599578063d7d9106b146105ba578063df5cf723146105cd578063e2d53d7c146105f4578063ef2f887b1461060757600080fd5b80639b6c9241116100ff5780639b6c924114610509578063acc8df9c1461055d578063b98d090814610566578063bafbcbe714610573578063cd22ed7b1461058657600080fd5b806372d18e8d146104b4578063886f1195146104c25780638b00ce7c146104d55780638da5cb5b146104e55780639b056bc1146104f657600080fd5b80635ac86ab7116101c95780636b532e9e1161018d5780636b532e9e1461043d5780636b92787e146104515780636d14a987146104645780636efb46361461048b578063715018a6146104ac57600080fd5b80635ac86ab71461037c5780635c155662146103af5780635c975abb146103cf5780635df45946146103d7578063683048351461041657600080fd5b80633563b0d1116102105780633563b0d11461030e578063416c7e5e1461032e5780634524c7e1146103415780634f739f7414610354578063595c6a671461037457600080fd5b806310d67a2f1461024d578063136439dd14610262578063171f1d5b146102755780631ad43189146102a45780632d89f6fc146102e0575b600080fd5b61026061025b366004614832565b610677565b005b61026061027036600461484f565b610733565b6102886102833660046149b9565b610872565b6040805192151583529015156020830152015b60405180910390f35b6102cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161029b565b6103006102ee366004614a27565b60cc6020526000908152604090205481565b60405190815260200161029b565b61032161031c366004614a9b565b6109fc565b60405161029b9190614bb6565b61026061033c366004614bd7565b610e94565b61026061034f36600461484f565b611009565b610367610362366004614c80565b611072565b60405161029b9190614d55565b610260611798565b61039f61038a366004614e1f565b606654600160ff9092169190911b9081161490565b604051901515815260200161029b565b6103c26103bd366004614e5f565b61185f565b60405161029b9190614f19565b606654610300565b6103fe7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161029b565b6103fe7f000000000000000000000000000000000000000000000000000000000000000081565b61026061044b366004614fe4565b50505050565b61026061045f366004615073565b611a27565b6103fe7f000000000000000000000000000000000000000000000000000000000000000081565b61049e6104993660046152db565b611bfd565b60405161029b92919061539b565b610260612b14565b60cb5463ffffffff166102cb565b6065546103fe906001600160a01b031681565b60cb546102cb9063ffffffff1681565b6033546001600160a01b03166103fe565b6102606105043660046153e4565b612b28565b61051c61051736600461548d565b612cf1565b60405161029b9190815163ffffffff908116825260208084015160ff1690830152604080840151821690830152606092830151169181019190915260800190565b61030060c95481565b60975461039f9060ff1681565b6102606105813660046154ce565b612e7d565b61026061059436600461484f565b61346a565b6105ac6105a73660046155d1565b613477565b60405161029b929190615608565b6103006105c8366004615629565b613609565b6103fe7f000000000000000000000000000000000000000000000000000000000000000081565b61039f610602366004614832565b61363a565b61030060ca5481565b61026061061e366004614832565b61376d565b7f00000000000000000000000000000000000000000000000000000000000000006102cb565b6102cb606481565b61026061065f366004615655565b6137e3565b61026061067236600461484f565b613904565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ee91906156b1565b6001600160a01b0316336001600160a01b0316146107275760405162461bcd60e51b815260040161071e906156ce565b60405180910390fd5b61073081613a60565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561077b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079f9190615718565b6107bb5760405162461bcd60e51b815260040161071e90615735565b606654818116146108345760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c6974790000000000000000606482015260840161071e565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106108ba576108ba61577d565b60200201518951600160200201518a602001516000600281106108df576108df61577d565b60200201518b602001516001600281106108fb576108fb61577d565b602090810291909101518c518d8301516040516109589a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c61097b9190615793565b90506109ee61099461098d8884613b57565b8690613bee565b61099c613c82565b6109e46109d5856109cf604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613b57565b6109de8c613d42565b90613bee565b886201d4c0613dd2565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a6291906156b1565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610aa4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac891906156b1565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b0a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b2e91906156b1565b9050600086516001600160401b03811115610b4b57610b4b614868565b604051908082528060200260200182016040528015610b7e57816020015b6060815260200190600190039081610b695790505b50905060005b8751811015610e86576000888281518110610ba157610ba161577d565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610c02573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c2a91908101906157b5565b905080516001600160401b03811115610c4557610c45614868565b604051908082528060200260200182016040528015610c9057816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610c635790505b50848481518110610ca357610ca361577d565b602002602001018190525060005b8151811015610e70576040518060600160405280876001600160a01b03166347b314e8858581518110610ce657610ce661577d565b60200260200101516040518263ffffffff1660e01b8152600401610d0c91815260200190565b602060405180830381865afa158015610d29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d4d91906156b1565b6001600160a01b03168152602001838381518110610d6d57610d6d61577d565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610d9b57610d9b61577d565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610df7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1b9190615845565b6001600160601b0316815250858581518110610e3957610e3961577d565b60200260200101518281518110610e5257610e5261577d565b60200260200101819052508080610e6890615884565b915050610cb1565b5050508080610e7e90615884565b915050610b84565b5093505050505b9392505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ef2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f1691906156b1565b6001600160a01b0316336001600160a01b031614610fc25760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a40161071e565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b611011613ff6565b606481111561106d5760405162461bcd60e51b815260206004820152602260248201527f7468726573686f6c642063616e2774206265206c6172676572207468616e2031604482015261030360f41b606482015260840161071e565b60c955565b61109d6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110dd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061110191906156b1565b905061112e6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061115e908b908990899060040161589f565b600060405180830381865afa15801561117b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111a391908101906158e9565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906111d5908b908b908b906004016159a0565b600060405180830381865afa1580156111f2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261121a91908101906158e9565b6040820152856001600160401b0381111561123757611237614868565b60405190808252806020026020018201604052801561126a57816020015b60608152602001906001900390816112555790505b50606082015260005b60ff81168711156116a9576000856001600160401b0381111561129857611298614868565b6040519080825280602002602001820160405280156112c1578160200160208202803683370190505b5083606001518360ff16815181106112db576112db61577d565b602002602001018190525060005b868110156115a95760008c6001600160a01b03166304ec63518a8a858181106113145761131461577d565b905060200201358e886000015186815181106113325761133261577d565b60200260200101516040518463ffffffff1660e01b815260040161136f9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561138c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113b091906159c9565b90506001600160c01b0381166114545760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a40161071e565b8a8a8560ff168181106114695761146961577d565b6001600160c01b03841692013560f81c9190911c60019081161415905061159657856001600160a01b031663dd9846b98a8a858181106114ab576114ab61577d565b905060200201358d8d8860ff168181106114c7576114c761577d565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa15801561151d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061154191906159f2565b85606001518560ff168151811061155a5761155a61577d565b602002602001015184815181106115735761157361577d565b63ffffffff909216602092830291909101909101528261159281615884565b9350505b50806115a181615884565b9150506112e9565b506000816001600160401b038111156115c4576115c4614868565b6040519080825280602002602001820160405280156115ed578160200160208202803683370190505b50905060005b8281101561166e5784606001518460ff16815181106116145761161461577d565b6020026020010151818151811061162d5761162d61577d565b60200260200101518282815181106116475761164761577d565b63ffffffff909216602092830291909101909101528061166681615884565b9150506115f3565b508084606001518460ff16815181106116895761168961577d565b6020026020010181905250505080806116a190615a0f565b915050611273565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061170e91906156b1565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611741908b908b908e90600401615a2f565b600060405180830381865afa15801561175e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261178691908101906158e9565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156117e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118049190615718565b6118205760405162461bcd60e51b815260040161071e90615735565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611891929190615a59565b600060405180830381865afa1580156118ae573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526118d691908101906158e9565b9050600084516001600160401b038111156118f3576118f3614868565b60405190808252806020026020018201604052801561191c578160200160208202803683370190505b50905060005b8551811015611a1d57866001600160a01b03166304ec635187838151811061194c5761194c61577d565b6020026020010151878685815181106119675761196761577d565b60200260200101516040518463ffffffff1660e01b81526004016119a49392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156119c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119e591906159c9565b6001600160c01b0316828281518110611a0057611a0061577d565b602090810291909101015280611a1581615884565b915050611922565b5095945050505050565b604051631619718360e21b81523360048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635865c60c906024016040805180830381865afa158015611a8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ab19190615aad565b8051909150611abf57600080fd5b611af6604051806080016040528060008152602001600063ffffffff16815260200160608152602001600063ffffffff1681525090565b85815263ffffffff438116602080840191909152908616606083015260408051601f860183900483028101830190915284815290859085908190840183828082843760009201919091525050505060408083019190915251611b5c908290602001615b36565b60408051601f19818403018152828252805160209182012060cb805463ffffffff908116600090815260cc90945293909220555416907f1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c4890611bbf908490615b36565b60405180910390a260cb54611bdb9063ffffffff166001615b89565b60cb805463ffffffff191663ffffffff92909216919091179055505050505050565b6040805180820190915260608082526020820152600084611c745760405162461bcd60e51b81526020600482015260376024820152600080516020615ed983398151915260448201527f7265733a20656d7074792071756f72756d20696e707574000000000000000000606482015260840161071e565b60408301515185148015611c8c575060a08301515185145b8015611c9c575060c08301515185145b8015611cac575060e08301515185145b611d165760405162461bcd60e51b81526020600482015260416024820152600080516020615ed983398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a40161071e565b82515160208401515114611d8e5760405162461bcd60e51b815260206004820152604460248201819052600080516020615ed9833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a40161071e565b4363ffffffff168463ffffffff1610611dfd5760405162461bcd60e51b815260206004820152603c6024820152600080516020615ed983398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b00000000606482015260840161071e565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611e3e57611e3e614868565b604051908082528060200260200182016040528015611e67578160200160208202803683370190505b506020820152866001600160401b03811115611e8557611e85614868565b604051908082528060200260200182016040528015611eae578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611ee257611ee2614868565b604051908082528060200260200182016040528015611f0b578160200160208202803683370190505b5081526020860151516001600160401b03811115611f2b57611f2b614868565b604051908082528060200260200182016040528015611f54578160200160208202803683370190505b50816020018190525060006120268a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611ffd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120219190615bb1565b614050565b905060005b8760200151518110156122c157612070886020015182815181106120515761205161577d565b6020026020010151805160009081526020918201519091526040902090565b836020015182815181106120865761208661577d565b602090810291909101015280156121465760208301516120a7600183615bce565b815181106120b7576120b761577d565b602002602001015160001c836020015182815181106120d8576120d861577d565b602002602001015160001c11612146576040805162461bcd60e51b8152602060048201526024810191909152600080516020615ed983398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f72746564606482015260840161071e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061218b5761218b61577d565b60200260200101518b8b6000015185815181106121aa576121aa61577d565b60200260200101516040518463ffffffff1660e01b81526004016121e79392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612204573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061222891906159c9565b6001600160c01b0316836000015182815181106122475761224761577d565b6020026020010181815250506122ad61098d61228184866000015185815181106122735761227361577d565b6020026020010151166140da565b8a6020015184815181106122975761229761577d565b602002602001015161410590919063ffffffff16565b9450806122b981615884565b91505061202b565b50506122cc836141e9565b60975490935060ff166000816122e3576000612365565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612341573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123659190615be5565b905060005b8a8110156129e35782156124c5578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106123c1576123c161577d565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612401573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124259190615be5565b61242f9190615bfe565b116124c55760405162461bcd60e51b81526020600482015260666024820152600080516020615ed983398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c40161071e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106125065761250661577d565b9050013560f81c60f81b60f81c8c8c60a00151858151811061252a5761252a61577d565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612586573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125aa9190615c16565b6001600160401b0319166125cd8a6040015183815181106120515761205161577d565b67ffffffffffffffff1916146126695760405162461bcd60e51b81526020600482015260616024820152600080516020615ed983398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c40161071e565b612699896040015182815181106126825761268261577d565b602002602001015187613bee90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106126dc576126dc61577d565b9050013560f81c60f81b60f81c8c8c60c0015185815181106127005761270061577d565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561275c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127809190615845565b856020015182815181106127965761279661577d565b6001600160601b039092166020928302919091018201528501518051829081106127c2576127c261577d565b6020026020010151856000015182815181106127e0576127e061577d565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156129ce576128588660000151828151811061282a5761282a61577d565b60200260200101518f8f868181106128445761284461577d565b600192013560f81c9290921c811614919050565b156129bc577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061289e5761289e61577d565b9050013560f81c60f81b60f81c8e896020015185815181106128c2576128c261577d565b60200260200101518f60e0015188815181106128e0576128e061577d565b602002602001015187815181106128f9576128f961577d565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa15801561295d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129819190615845565b87518051859081106129955761299561577d565b602002602001018181516129a99190615c41565b6001600160601b03169052506001909101905b806129c681615884565b915050612804565b505080806129db90615884565b91505061236a565b5050506000806129fd8c868a606001518b60800151610872565b9150915081612a6e5760405162461bcd60e51b81526020600482015260436024820152600080516020615ed983398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a40161071e565b80612acf5760405162461bcd60e51b81526020600482015260396024820152600080516020615ed983398151915260448201527f7265733a207369676e617475726520697320696e76616c696400000000000000606482015260840161071e565b50506000878260200151604051602001612aea929190615c69565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612b1c613ff6565b612b266000614284565b565b604051631619718360e21b81523360048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635865c60c906024016040805180830381865afa158015612b8e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bb29190615aad565b8051909150612bc057600080fd5b6040805160a0810182526060818301819052600090820181905260808201524363ffffffff16815260208082018990528251601f8701829004820281018201909352858352909190869086908190840183828082843760009201919091525050505060408083019190915263ffffffff8716606083015260ff8416608083015251612c4f908290602001615cb1565b60408051601f19818403018152828252805160209182012060cb805463ffffffff908116600090815260cc90945293909220555416907f903a8215cb7ce14c9230c13623ee3c35bc42c4a0cb4fed2ff3a24ca32c7773f390612cb2908490615cb1565b60405180910390a260cb54612cce9063ffffffff166001615b89565b60cb805463ffffffff191663ffffffff9290921691909117905550505050505050565b604080516080810182526000808252602082018190528183018190526060820152905160ce90612d249085908590615d22565b9081526040519081900360200190205463ffffffff6501000000000090910416612d905760405162461bcd60e51b815260206004820152601f60248201527f5265717565737465642066656564206973206e6f7420737570706f7274656400604482015260640161071e565b600060ce8484604051612da4929190615d22565b9081526040805191829003602090810183206080840183525463ffffffff808216855260ff640100000000830416928501929092526501000000000081048216928401839052600160481b9004166060830152909150606490612e079043615d32565b63ffffffff1610612e745760405162461bcd60e51b815260206004820152603160248201527f526571756573746564206665656420686173206e6f74206265656e207570646160448201527074656420696e2031303020626c6f636b7360781b606482015260840161071e565b90505b92915050565b604051631619718360e21b81523360048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635865c60c906024016040805180830381865afa158015612ee3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f079190615aad565b8051909150612f1557600080fd5b6000612f246020870187614a27565b9050366000612f366040890189615d4f565b90925090506000612f4d60808a0160608b01614a27565b905086612fa85760405162461bcd60e51b8152602060048201526024808201527f41676772656761746f72206d757374207375626d6974207461736b20726573706044820152636f6e736560e01b606482015260840161071e565b855187146130165760405162461bcd60e51b815260206004820152603560248201527f726573706f6e7365206c656e677468206d757374206d617463682061676772656044820152740cec2e8ca40e6d2cedcc2e8eae4ca40d8cadccee8d605b1b606482015260840161071e565b60cd60008989600081811061302d5761302d61577d565b905060200281019061303f9190615d95565b613050906060810190604001614a27565b63ffffffff168152602081019190915260400160002054156130c95760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b606482015260840161071e565b6130f37f000000000000000000000000000000000000000000000000000000000000000085615b89565b63ffffffff164363ffffffff1611156131645760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b606482015260840161071e565b60005b878110156132e8576040805180820190915263ffffffff4316815260006020820181905260cd908b8b858181106131a0576131a061577d565b90506020028101906131b29190615d95565b6131c3906060810190604001614a27565b63ffffffff1663ffffffff1681526020019081526020016000208a8a848181106131ef576131ef61577d565b90506020028101906132019190615d95565b82604051602001613213929190615db5565b60408051601f198184030181529190528051602091820120825460018101845560009384529190922001558989838181106132505761325061577d565b90506020028101906132629190615d95565b613273906060810190604001614a27565b63ffffffff167fbf877b433f8d349e650f5f2d05288ff1b727ac7652853301d6c73558b2757f7f8b8b858181106132ac576132ac61577d565b90506020028101906132be9190615d95565b836040516132cd929190615db5565b60405180910390a250806132e081615884565b915050613167565b506040518060800160405280898960008181106133075761330761577d565b90506020028101906133199190615d95565b613327906020810190614a27565b63ffffffff168152602001898960008181106133455761334561577d565b90506020028101906133579190615d95565b613368906040810190602001614a27565b60ff1681526020014363ffffffff1681526020018989600081811061338f5761338f61577d565b90506020028101906133a19190615d95565b6133af906020810190614a27565b63ffffffff16905260ce6133c660208c018c615d4f565b6040516133d4929190615d22565b908152604080516020928190038301902083518154938501519285015160609095015163ffffffff91821664ffffffffff199095169490941764010000000060ff90941693909302929092176cffffffffffffffff0000000000191665010000000000948316949094026cffffffff000000000000000000191693909317600160481b9190921602179055505050505050505050565b613472613ff6565b60ca55565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106134b2576134b261577d565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906134ee9088908690600401615a59565b600060405180830381865afa15801561350b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261353391908101906158e9565b6000815181106135455761354561577d565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156135b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135d591906159c9565b6001600160c01b0316905060006135eb826142d6565b9050816135f98a838a6109fc565b9550955050505050935093915050565b60cd602052816000526040600020818154811061362557600080fd5b90600052602060002001600091509150505481565b604051631619718360e21b815233600482015260009081906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635865c60c906024016040805180830381865afa1580156136a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136c69190615aad565b80519091506136d457600080fd5b604051631619718360e21b81526001600160a01b0384811660048301526000917f000000000000000000000000000000000000000000000000000000000000000090911690635865c60c906024016040805180830381865afa15801561373e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137629190615aad565b511515949350505050565b613775613ff6565b6001600160a01b0381166137da5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161071e565b61073081614284565b600054610100900460ff16158080156138035750600054600160ff909116105b8061381d5750303b15801561381d575060005460ff166001145b6138805760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161071e565b6000805460ff1916600117905580156138a3576000805461ff0019166101001790555b6138ae8560006143a2565b6138b784614284565b80156138fd576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613957573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061397b91906156b1565b6001600160a01b0316336001600160a01b0316146139ab5760405162461bcd60e51b815260040161071e906156ce565b606654198119606654191614613a295760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c6974790000000000000000606482015260840161071e565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610867565b6001600160a01b038116613aee5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a40161071e565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152613b73614743565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015613ba657613ba8565bfe5b5080613be65760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b604482015260640161071e565b505092915050565b6040805180820190915260008082526020820152613c0a614761565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015613ba6575080613be65760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b604482015260640161071e565b613c8a61477f565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613d72600080516020615eb983398151915286615793565b90505b613d7e8161448c565b9093509150600080516020615eb9833981519152828309831415613db8576040805180820190915290815260208101919091529392505050565b600080516020615eb9833981519152600182089050613d75565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613e046147a4565b60005b6002811015613fc9576000613e1d826006615e77565b9050848260028110613e3157613e3161577d565b60200201515183613e43836000615bfe565b600c8110613e5357613e5361577d565b6020020152848260028110613e6a57613e6a61577d565b60200201516020015183826001613e819190615bfe565b600c8110613e9157613e9161577d565b6020020152838260028110613ea857613ea861577d565b6020020151515183613ebb836002615bfe565b600c8110613ecb57613ecb61577d565b6020020152838260028110613ee257613ee261577d565b6020020151516001602002015183613efb836003615bfe565b600c8110613f0b57613f0b61577d565b6020020152838260028110613f2257613f2261577d565b602002015160200151600060028110613f3d57613f3d61577d565b602002015183613f4e836004615bfe565b600c8110613f5e57613f5e61577d565b6020020152838260028110613f7557613f7561577d565b602002015160200151600160028110613f9057613f9061577d565b602002015183613fa1836005615bfe565b600c8110613fb157613fb161577d565b60200201525080613fc181615884565b915050613e07565b50613fd26147c3565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6033546001600160a01b03163314612b265760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161071e565b60008061405c8461450e565b9050808360ff166001901b11612e745760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c756500606482015260840161071e565b6000805b8215612e77576140ef600184615bce565b90921691806140fd81615e96565b9150506140de565b60408051808201909152600080825260208201526102008261ffff16106141615760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b604482015260640161071e565b8161ffff1660011415614175575081612e77565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff16106141de57600161ffff871660ff83161c811614156141c1576141be8484613bee565b93505b6141cb8384613bee565b92506201fffe600192831b169101614191565b509195945050505050565b6040805180820190915260008082526020820152815115801561420e57506020820151155b1561422c575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615eb9833981519152846020015161425f9190615793565b61427790600080516020615eb9833981519152615bce565b905292915050565b919050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606000806142e4846140da565b61ffff166001600160401b038111156142ff576142ff614868565b6040519080825280601f01601f191660200182016040528015614329576020820181803683370190505b5090506000805b825182108015614341575061010081105b15614398576001811b935085841615614388578060f81b83838151811061436a5761436a61577d565b60200101906001600160f81b031916908160001a9053508160010191505b61439181615884565b9050614330565b5090949350505050565b6065546001600160a01b03161580156143c357506001600160a01b03821615155b6144455760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a40161071e565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261448882613a60565b5050565b60008080600080516020615eb98339815191526003600080516020615eb983398151915286600080516020615eb9833981519152888909090890506000614502827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615eb983398151915261469b565b91959194509092505050565b6000610100825111156145975760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a40161071e565b81516145a557506000919050565b600080836000815181106145bb576145bb61577d565b0160200151600160f89190911c81901b92505b8451811015614692578481815181106145e9576145e961577d565b0160200151600160f89190911c1b915082821161467e5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a40161071e565b9181179161468b81615884565b90506145ce565b50909392505050565b6000806146a66147c3565b6146ae6147e1565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613ba65750826147385760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c757265000000000000604482015260640161071e565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806147926147ff565b815260200161479f6147ff565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461073057600080fd5b60006020828403121561484457600080fd5b8135612e748161481d565b60006020828403121561486157600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156148a0576148a0614868565b60405290565b60405161010081016001600160401b03811182821017156148a0576148a0614868565b604051601f8201601f191681016001600160401b03811182821017156148f1576148f1614868565b604052919050565b60006040828403121561490b57600080fd5b61491361487e565b9050813581526020820135602082015292915050565b600082601f83011261493a57600080fd5b61494261487e565b80604084018581111561495457600080fd5b845b8181101561496e578035845260209384019301614956565b509095945050505050565b60006080828403121561498b57600080fd5b61499361487e565b905061499f8383614929565b81526149ae8360408401614929565b602082015292915050565b60008060008061012085870312156149d057600080fd5b843593506149e186602087016148f9565b92506149f08660608701614979565b91506149ff8660e087016148f9565b905092959194509250565b63ffffffff8116811461073057600080fd5b803561427f81614a0a565b600060208284031215614a3957600080fd5b8135612e7481614a0a565b60006001600160401b03831115614a5d57614a5d614868565b614a70601f8401601f19166020016148c9565b9050828152838383011115614a8457600080fd5b828260208301376000602084830101529392505050565b600080600060608486031215614ab057600080fd5b8335614abb8161481d565b925060208401356001600160401b03811115614ad657600080fd5b8401601f81018613614ae757600080fd5b614af686823560208401614a44565b9250506040840135614b0781614a0a565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614ba8578385038a52825180518087529087019087870190845b81811015614b9357835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614b4f565b50509a87019a95505091850191600101614b31565b509298975050505050505050565b602081526000610e8d6020830184614b12565b801515811461073057600080fd5b600060208284031215614be957600080fd5b8135612e7481614bc9565b60008083601f840112614c0657600080fd5b5081356001600160401b03811115614c1d57600080fd5b602083019150836020828501011115614c3557600080fd5b9250929050565b60008083601f840112614c4e57600080fd5b5081356001600160401b03811115614c6557600080fd5b6020830191508360208260051b8501011115614c3557600080fd5b60008060008060008060808789031215614c9957600080fd5b8635614ca48161481d565b95506020870135614cb481614a0a565b945060408701356001600160401b0380821115614cd057600080fd5b614cdc8a838b01614bf4565b90965094506060890135915080821115614cf557600080fd5b50614d0289828a01614c3c565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b83811015614d4a57815163ffffffff1687529582019590820190600101614d28565b509495945050505050565b600060208083528351608082850152614d7160a0850182614d14565b905081850151601f1980868403016040870152614d8e8383614d14565b92506040870151915080868403016060870152614dab8383614d14565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614e025784878303018452614df0828751614d14565b95880195938801939150600101614dd6565b509998505050505050505050565b60ff8116811461073057600080fd5b600060208284031215614e3157600080fd5b8135612e7481614e10565b60006001600160401b03821115614e5557614e55614868565b5060051b60200190565b600080600060608486031215614e7457600080fd5b8335614e7f8161481d565b92506020848101356001600160401b03811115614e9b57600080fd5b8501601f81018713614eac57600080fd5b8035614ebf614eba82614e3c565b6148c9565b81815260059190911b82018301908381019089831115614ede57600080fd5b928401925b82841015614efc57833582529284019290840190614ee3565b8096505050505050614f1060408501614a1c565b90509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614f5157835183529284019291840191600101614f35565b50909695505050505050565b600060408284031215614f6f57600080fd5b50919050565b600082601f830112614f8657600080fd5b81356020614f96614eba83614e3c565b82815260069290921b84018101918181019086841115614fb557600080fd5b8286015b84811015614fd957614fcb88826148f9565b835291830191604001614fb9565b509695505050505050565b60008060008060c08587031215614ffa57600080fd5b84356001600160401b038082111561501157600080fd5b908601906080828903121561502557600080fd5b8195506150358860208901614f5d565b94506150448860608901614f5d565b935060a087013591508082111561505a57600080fd5b5061506787828801614f75565b91505092959194509250565b6000806000806060858703121561508957600080fd5b84359350602085013561509b81614a0a565b925060408501356001600160401b038111156150b657600080fd5b6150c287828801614bf4565b95989497509550505050565b600082601f8301126150df57600080fd5b813560206150ef614eba83614e3c565b82815260059290921b8401810191818101908684111561510e57600080fd5b8286015b84811015614fd957803561512581614a0a565b8352918301918301615112565b600082601f83011261514357600080fd5b81356020615153614eba83614e3c565b82815260059290921b8401810191818101908684111561517257600080fd5b8286015b84811015614fd95780356001600160401b038111156151955760008081fd5b6151a38986838b01016150ce565b845250918301918301615176565b600061018082840312156151c457600080fd5b6151cc6148a6565b905081356001600160401b03808211156151e557600080fd5b6151f1858386016150ce565b8352602084013591508082111561520757600080fd5b61521385838601614f75565b6020840152604084013591508082111561522c57600080fd5b61523885838601614f75565b604084015261524a8560608601614979565b606084015261525c8560e086016148f9565b608084015261012084013591508082111561527657600080fd5b615282858386016150ce565b60a084015261014084013591508082111561529c57600080fd5b6152a8858386016150ce565b60c08401526101608401359150808211156152c257600080fd5b506152cf84828501615132565b60e08301525092915050565b6000806000806000608086880312156152f357600080fd5b8535945060208601356001600160401b038082111561531157600080fd5b61531d89838a01614bf4565b90965094506040880135915061533282614a0a565b9092506060870135908082111561534857600080fd5b50615355888289016151b1565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614d4a5781516001600160601b031687529582019590820190600101615376565b60408152600083516040808401526153b66080840182615362565b90506020850151603f198483030160608501526153d38282615362565b925050508260208301529392505050565b6000806000806000608086880312156153fc57600080fd5b85356001600160401b038082111561541357600080fd5b818801915088601f83011261542757600080fd5b61543689833560208501614a44565b96506020880135915061544882614a0a565b9094506040870135908082111561545e57600080fd5b5061546b88828901614bf4565b909450925050606086013561547f81614e10565b809150509295509295909350565b600080602083850312156154a057600080fd5b82356001600160401b038111156154b657600080fd5b6154c285828601614bf4565b90969095509350505050565b600080600080606085870312156154e457600080fd5b84356001600160401b03808211156154fb57600080fd5b9086019060a0828903121561550f57600080fd5b909450602090868201358181111561552657600080fd5b61553289828a01614c3c565b90965094505060408701358181111561554a57600080fd5b8701601f8101891361555b57600080fd5b8035615569614eba82614e3c565b81815260059190911b8201840190848101908b83111561558857600080fd5b8584015b838110156155c0578035868111156155a45760008081fd5b6155b28e89838901016151b1565b84525091860191860161558c565b50989b979a50959850505050505050565b6000806000606084860312156155e657600080fd5b83356155f18161481d565b9250602084013591506040840135614b0781614a0a565b8281526040602082015260006156216040830184614b12565b949350505050565b6000806040838503121561563c57600080fd5b823561564781614a0a565b946020939093013593505050565b6000806000806080858703121561566b57600080fd5b84356156768161481d565b935060208501356156868161481d565b925060408501356156968161481d565b915060608501356156a68161481d565b939692955090935050565b6000602082840312156156c357600080fd5b8151612e748161481d565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561572a57600080fd5b8151612e7481614bc9565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000826157b057634e487b7160e01b600052601260045260246000fd5b500690565b600060208083850312156157c857600080fd5b82516001600160401b038111156157de57600080fd5b8301601f810185136157ef57600080fd5b80516157fd614eba82614e3c565b81815260059190911b8201830190838101908783111561581c57600080fd5b928401925b8284101561583a57835182529284019290840190615821565b979650505050505050565b60006020828403121561585757600080fd5b81516001600160601b0381168114612e7457600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156158985761589861586e565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156158cc57600080fd5b8260051b8085606085013760009201606001918252509392505050565b600060208083850312156158fc57600080fd5b82516001600160401b0381111561591257600080fd5b8301601f8101851361592357600080fd5b8051615931614eba82614e3c565b81815260059190911b8201830190838101908783111561595057600080fd5b928401925b8284101561583a57835161596881614a0a565b82529284019290840190615955565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006159c0604083018486615977565b95945050505050565b6000602082840312156159db57600080fd5b81516001600160c01b0381168114612e7457600080fd5b600060208284031215615a0457600080fd5b8151612e7481614a0a565b600060ff821660ff811415615a2657615a2661586e565b60010192915050565b604081526000615a43604083018587615977565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b81811015615aa057845183529383019391830191600101615a84565b5090979650505050505050565b600060408284031215615abf57600080fd5b615ac761487e565b82518152602083015160038110615add57600080fd5b60208201529392505050565b6000815180845260005b81811015615b0f57602081850181015186830182015201615af3565b81811115615b21576000602083870101525b50601f01601f19169290920160200192915050565b60208152815160208201526000602083015163ffffffff80821660408501526040850151915060806060850152615b7060a0850183615ae9565b9150806060860151166080850152508091505092915050565b600063ffffffff808316818516808303821115615ba857615ba861586e565b01949350505050565b600060208284031215615bc357600080fd5b8151612e7481614e10565b600082821015615be057615be061586e565b500390565b600060208284031215615bf757600080fd5b5051919050565b60008219821115615c1157615c1161586e565b500190565b600060208284031215615c2857600080fd5b815167ffffffffffffffff1981168114612e7457600080fd5b60006001600160601b0383811690831681811015615c6157615c6161586e565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615ca457815185529382019390820190600101615c88565b5092979650505050505050565b60208152600063ffffffff808451166020840152602084015160a06040850152615cde60c0850182615ae9565b90506040850151601f19858303016060860152615cfb8282615ae9565b91505081606086015116608085015260ff60808601511660a0850152809250505092915050565b8183823760009101908152919050565b600063ffffffff83811690831681811015615c6157615c6161586e565b6000808335601e19843603018112615d6657600080fd5b8301803591506001600160401b03821115615d8057600080fd5b602001915036819003821315614c3557600080fd5b60008235607e19833603018112615dab57600080fd5b9190910192915050565b6060815260008335615dc681614a0a565b63ffffffff9081166060840152602085013590615de282614a0a565b9081166080840152604085013590615df982614a0a565b1660a0830152606084013536859003601e19018112615e1757600080fd5b840180356001600160401b03811115615e2f57600080fd5b803603861315615e3e57600080fd5b608060c0850152615e5660e085018260208501615977565b92505050610e8d6020830184805163ffffffff168252602090810151910152565b6000816000190483118215151615615e9157615e9161586e565b500290565b600061ffff80831681811415615eae57615eae61586e565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220737770de2c959c22f5aa3c648a5fa27d7ffe1e00037714b7bf8ea6442bb47a4164736f6c634300080c0033",
}

// ContractPriceAggregatorTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractPriceAggregatorTaskManagerMetaData.ABI instead.
var ContractPriceAggregatorTaskManagerABI = ContractPriceAggregatorTaskManagerMetaData.ABI

// ContractPriceAggregatorTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractPriceAggregatorTaskManagerMetaData.Bin instead.
var ContractPriceAggregatorTaskManagerBin = ContractPriceAggregatorTaskManagerMetaData.Bin

// DeployContractPriceAggregatorTaskManager deploys a new Ethereum contract, binding an instance of ContractPriceAggregatorTaskManager to it.
func DeployContractPriceAggregatorTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractPriceAggregatorTaskManager, error) {
	parsed, err := ContractPriceAggregatorTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractPriceAggregatorTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractPriceAggregatorTaskManager{ContractPriceAggregatorTaskManagerCaller: ContractPriceAggregatorTaskManagerCaller{contract: contract}, ContractPriceAggregatorTaskManagerTransactor: ContractPriceAggregatorTaskManagerTransactor{contract: contract}, ContractPriceAggregatorTaskManagerFilterer: ContractPriceAggregatorTaskManagerFilterer{contract: contract}}, nil
}

// ContractPriceAggregatorTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractPriceAggregatorTaskManager struct {
	ContractPriceAggregatorTaskManagerCaller     // Read-only binding to the contract
	ContractPriceAggregatorTaskManagerTransactor // Write-only binding to the contract
	ContractPriceAggregatorTaskManagerFilterer   // Log filterer for contract events
}

// ContractPriceAggregatorTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractPriceAggregatorTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPriceAggregatorTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractPriceAggregatorTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPriceAggregatorTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractPriceAggregatorTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPriceAggregatorTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractPriceAggregatorTaskManagerSession struct {
	Contract     *ContractPriceAggregatorTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ContractPriceAggregatorTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractPriceAggregatorTaskManagerCallerSession struct {
	Contract *ContractPriceAggregatorTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// ContractPriceAggregatorTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractPriceAggregatorTaskManagerTransactorSession struct {
	Contract     *ContractPriceAggregatorTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// ContractPriceAggregatorTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractPriceAggregatorTaskManagerRaw struct {
	Contract *ContractPriceAggregatorTaskManager // Generic contract binding to access the raw methods on
}

// ContractPriceAggregatorTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractPriceAggregatorTaskManagerCallerRaw struct {
	Contract *ContractPriceAggregatorTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractPriceAggregatorTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractPriceAggregatorTaskManagerTransactorRaw struct {
	Contract *ContractPriceAggregatorTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractPriceAggregatorTaskManager creates a new instance of ContractPriceAggregatorTaskManager, bound to a specific deployed contract.
func NewContractPriceAggregatorTaskManager(address common.Address, backend bind.ContractBackend) (*ContractPriceAggregatorTaskManager, error) {
	contract, err := bindContractPriceAggregatorTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManager{ContractPriceAggregatorTaskManagerCaller: ContractPriceAggregatorTaskManagerCaller{contract: contract}, ContractPriceAggregatorTaskManagerTransactor: ContractPriceAggregatorTaskManagerTransactor{contract: contract}, ContractPriceAggregatorTaskManagerFilterer: ContractPriceAggregatorTaskManagerFilterer{contract: contract}}, nil
}

// NewContractPriceAggregatorTaskManagerCaller creates a new read-only instance of ContractPriceAggregatorTaskManager, bound to a specific deployed contract.
func NewContractPriceAggregatorTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractPriceAggregatorTaskManagerCaller, error) {
	contract, err := bindContractPriceAggregatorTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerCaller{contract: contract}, nil
}

// NewContractPriceAggregatorTaskManagerTransactor creates a new write-only instance of ContractPriceAggregatorTaskManager, bound to a specific deployed contract.
func NewContractPriceAggregatorTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractPriceAggregatorTaskManagerTransactor, error) {
	contract, err := bindContractPriceAggregatorTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerTransactor{contract: contract}, nil
}

// NewContractPriceAggregatorTaskManagerFilterer creates a new log filterer instance of ContractPriceAggregatorTaskManager, bound to a specific deployed contract.
func NewContractPriceAggregatorTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractPriceAggregatorTaskManagerFilterer, error) {
	contract, err := bindContractPriceAggregatorTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerFilterer{contract: contract}, nil
}

// bindContractPriceAggregatorTaskManager binds a generic wrapper to an already deployed contract.
func bindContractPriceAggregatorTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractPriceAggregatorTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractPriceAggregatorTaskManager.Contract.ContractPriceAggregatorTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.ContractPriceAggregatorTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.ContractPriceAggregatorTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractPriceAggregatorTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTQUORUMTHRESHOLD is a free data retrieval call binding the contract method 0xacc8df9c.
//
// Solidity: function DEFAULT_QUORUM_THRESHOLD() view returns(uint256)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) DEFAULTQUORUMTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "DEFAULT_QUORUM_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTQUORUMTHRESHOLD is a free data retrieval call binding the contract method 0xacc8df9c.
//
// Solidity: function DEFAULT_QUORUM_THRESHOLD() view returns(uint256)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) DEFAULTQUORUMTHRESHOLD() (*big.Int, error) {
	return _ContractPriceAggregatorTaskManager.Contract.DEFAULTQUORUMTHRESHOLD(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// DEFAULTQUORUMTHRESHOLD is a free data retrieval call binding the contract method 0xacc8df9c.
//
// Solidity: function DEFAULT_QUORUM_THRESHOLD() view returns(uint256)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) DEFAULTQUORUMTHRESHOLD() (*big.Int, error) {
	return _ContractPriceAggregatorTaskManager.Contract.DEFAULTQUORUMTHRESHOLD(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// DEFAULTREQUIREDNUMOFSOURCES is a free data retrieval call binding the contract method 0xef2f887b.
//
// Solidity: function DEFAULT_REQUIRED_NUM_OF_SOURCES() view returns(uint256)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) DEFAULTREQUIREDNUMOFSOURCES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "DEFAULT_REQUIRED_NUM_OF_SOURCES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTREQUIREDNUMOFSOURCES is a free data retrieval call binding the contract method 0xef2f887b.
//
// Solidity: function DEFAULT_REQUIRED_NUM_OF_SOURCES() view returns(uint256)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) DEFAULTREQUIREDNUMOFSOURCES() (*big.Int, error) {
	return _ContractPriceAggregatorTaskManager.Contract.DEFAULTREQUIREDNUMOFSOURCES(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// DEFAULTREQUIREDNUMOFSOURCES is a free data retrieval call binding the contract method 0xef2f887b.
//
// Solidity: function DEFAULT_REQUIRED_NUM_OF_SOURCES() view returns(uint256)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) DEFAULTREQUIREDNUMOFSOURCES() (*big.Int, error) {
	return _ContractPriceAggregatorTaskManager.Contract.DEFAULTREQUIREDNUMOFSOURCES(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractPriceAggregatorTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractPriceAggregatorTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractPriceAggregatorTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractPriceAggregatorTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractPriceAggregatorTaskManager.Contract.AllTaskHashes(&_ContractPriceAggregatorTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractPriceAggregatorTaskManager.Contract.AllTaskHashes(&_ContractPriceAggregatorTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0xd7d9106b.
//
// Solidity: function allTaskResponses(uint32 , uint256 ) view returns(bytes32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0xd7d9106b.
//
// Solidity: function allTaskResponses(uint32 , uint256 ) view returns(bytes32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) AllTaskResponses(arg0 uint32, arg1 *big.Int) ([32]byte, error) {
	return _ContractPriceAggregatorTaskManager.Contract.AllTaskResponses(&_ContractPriceAggregatorTaskManager.CallOpts, arg0, arg1)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0xd7d9106b.
//
// Solidity: function allTaskResponses(uint32 , uint256 ) view returns(bytes32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) AllTaskResponses(arg0 uint32, arg1 *big.Int) ([32]byte, error) {
	return _ContractPriceAggregatorTaskManager.Contract.AllTaskResponses(&_ContractPriceAggregatorTaskManager.CallOpts, arg0, arg1)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractPriceAggregatorTaskManager.Contract.BlsApkRegistry(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractPriceAggregatorTaskManager.Contract.BlsApkRegistry(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractPriceAggregatorTaskManager.Contract.CheckSignatures(&_ContractPriceAggregatorTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractPriceAggregatorTaskManager.Contract.CheckSignatures(&_ContractPriceAggregatorTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Delegation(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Delegation(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// FetchLatestAggregatedPrice is a free data retrieval call binding the contract method 0x9b6c9241.
//
// Solidity: function fetchLatestAggregatedPrice(string feedName) view returns((uint32,uint8,uint32,uint32))
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) FetchLatestAggregatedPrice(opts *bind.CallOpts, feedName string) (IPriceAggregatorTaskManagerAggregatedPrice, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "fetchLatestAggregatedPrice", feedName)

	if err != nil {
		return *new(IPriceAggregatorTaskManagerAggregatedPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(IPriceAggregatorTaskManagerAggregatedPrice)).(*IPriceAggregatorTaskManagerAggregatedPrice)

	return out0, err

}

// FetchLatestAggregatedPrice is a free data retrieval call binding the contract method 0x9b6c9241.
//
// Solidity: function fetchLatestAggregatedPrice(string feedName) view returns((uint32,uint8,uint32,uint32))
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) FetchLatestAggregatedPrice(feedName string) (IPriceAggregatorTaskManagerAggregatedPrice, error) {
	return _ContractPriceAggregatorTaskManager.Contract.FetchLatestAggregatedPrice(&_ContractPriceAggregatorTaskManager.CallOpts, feedName)
}

// FetchLatestAggregatedPrice is a free data retrieval call binding the contract method 0x9b6c9241.
//
// Solidity: function fetchLatestAggregatedPrice(string feedName) view returns((uint32,uint8,uint32,uint32))
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) FetchLatestAggregatedPrice(feedName string) (IPriceAggregatorTaskManagerAggregatedPrice, error) {
	return _ContractPriceAggregatorTaskManager.Contract.FetchLatestAggregatedPrice(&_ContractPriceAggregatorTaskManager.CallOpts, feedName)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractPriceAggregatorTaskManager.Contract.GetCheckSignaturesIndices(&_ContractPriceAggregatorTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractPriceAggregatorTaskManager.Contract.GetCheckSignaturesIndices(&_ContractPriceAggregatorTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractPriceAggregatorTaskManager.Contract.GetOperatorState(&_ContractPriceAggregatorTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractPriceAggregatorTaskManager.Contract.GetOperatorState(&_ContractPriceAggregatorTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractPriceAggregatorTaskManager.Contract.GetOperatorState0(&_ContractPriceAggregatorTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractPriceAggregatorTaskManager.Contract.GetOperatorState0(&_ContractPriceAggregatorTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractPriceAggregatorTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractPriceAggregatorTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractPriceAggregatorTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractPriceAggregatorTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractPriceAggregatorTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractPriceAggregatorTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// IsValidOperator is a free data retrieval call binding the contract method 0xe2d53d7c.
//
// Solidity: function isValidOperator(address operatorAddress) view returns(bool)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) IsValidOperator(opts *bind.CallOpts, operatorAddress common.Address) (bool, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "isValidOperator", operatorAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidOperator is a free data retrieval call binding the contract method 0xe2d53d7c.
//
// Solidity: function isValidOperator(address operatorAddress) view returns(bool)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) IsValidOperator(operatorAddress common.Address) (bool, error) {
	return _ContractPriceAggregatorTaskManager.Contract.IsValidOperator(&_ContractPriceAggregatorTaskManager.CallOpts, operatorAddress)
}

// IsValidOperator is a free data retrieval call binding the contract method 0xe2d53d7c.
//
// Solidity: function isValidOperator(address operatorAddress) view returns(bool)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) IsValidOperator(operatorAddress common.Address) (bool, error) {
	return _ContractPriceAggregatorTaskManager.Contract.IsValidOperator(&_ContractPriceAggregatorTaskManager.CallOpts, operatorAddress)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractPriceAggregatorTaskManager.Contract.LatestTaskNum(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractPriceAggregatorTaskManager.Contract.LatestTaskNum(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) Owner() (common.Address, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Owner(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Owner(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Paused(&_ContractPriceAggregatorTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Paused(&_ContractPriceAggregatorTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Paused0(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Paused0(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractPriceAggregatorTaskManager.Contract.PauserRegistry(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractPriceAggregatorTaskManager.Contract.PauserRegistry(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractPriceAggregatorTaskManager.Contract.RegistryCoordinator(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractPriceAggregatorTaskManager.Contract.RegistryCoordinator(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractPriceAggregatorTaskManager.Contract.StakeRegistry(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractPriceAggregatorTaskManager.Contract.StakeRegistry(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractPriceAggregatorTaskManager.Contract.StaleStakesForbidden(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractPriceAggregatorTaskManager.Contract.StaleStakesForbidden(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractPriceAggregatorTaskManager.Contract.TaskNumber(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractPriceAggregatorTaskManager.Contract.TaskNumber(&_ContractPriceAggregatorTaskManager.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractPriceAggregatorTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractPriceAggregatorTaskManager.Contract.TrySignatureAndApkVerification(&_ContractPriceAggregatorTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractPriceAggregatorTaskManager.Contract.TrySignatureAndApkVerification(&_ContractPriceAggregatorTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x6b92787e.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, numberToBeSquared *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "createNewTask", numberToBeSquared, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x6b92787e.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) CreateNewTask(numberToBeSquared *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.CreateNewTask(&_ContractPriceAggregatorTaskManager.TransactOpts, numberToBeSquared, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x6b92787e.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) CreateNewTask(numberToBeSquared *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.CreateNewTask(&_ContractPriceAggregatorTaskManager.TransactOpts, numberToBeSquared, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Initialize(&_ContractPriceAggregatorTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Initialize(&_ContractPriceAggregatorTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Pause(&_ContractPriceAggregatorTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Pause(&_ContractPriceAggregatorTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.PauseAll(&_ContractPriceAggregatorTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.PauseAll(&_ContractPriceAggregatorTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x6b532e9e.
//
// Solidity: function raiseAndResolveChallenge((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IPriceAggregatorTaskManagerTask, taskResponse IPriceAggregatorTaskManagerTaskResponse, taskResponseMetadata IPriceAggregatorTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x6b532e9e.
//
// Solidity: function raiseAndResolveChallenge((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) RaiseAndResolveChallenge(task IPriceAggregatorTaskManagerTask, taskResponse IPriceAggregatorTaskManagerTaskResponse, taskResponseMetadata IPriceAggregatorTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.RaiseAndResolveChallenge(&_ContractPriceAggregatorTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x6b532e9e.
//
// Solidity: function raiseAndResolveChallenge((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) RaiseAndResolveChallenge(task IPriceAggregatorTaskManagerTask, taskResponse IPriceAggregatorTaskManagerTaskResponse, taskResponseMetadata IPriceAggregatorTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.RaiseAndResolveChallenge(&_ContractPriceAggregatorTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.RenounceOwnership(&_ContractPriceAggregatorTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.RenounceOwnership(&_ContractPriceAggregatorTaskManager.TransactOpts)
}

// RequestPriceFeedUpdate is a paid mutator transaction binding the contract method 0x9b056bc1.
//
// Solidity: function requestPriceFeedUpdate(string feedName, uint32 quorumThresholdPercentage, bytes quorumNumbers, uint8 minNumOfOracleNetworks) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) RequestPriceFeedUpdate(opts *bind.TransactOpts, feedName string, quorumThresholdPercentage uint32, quorumNumbers []byte, minNumOfOracleNetworks uint8) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "requestPriceFeedUpdate", feedName, quorumThresholdPercentage, quorumNumbers, minNumOfOracleNetworks)
}

// RequestPriceFeedUpdate is a paid mutator transaction binding the contract method 0x9b056bc1.
//
// Solidity: function requestPriceFeedUpdate(string feedName, uint32 quorumThresholdPercentage, bytes quorumNumbers, uint8 minNumOfOracleNetworks) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) RequestPriceFeedUpdate(feedName string, quorumThresholdPercentage uint32, quorumNumbers []byte, minNumOfOracleNetworks uint8) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.RequestPriceFeedUpdate(&_ContractPriceAggregatorTaskManager.TransactOpts, feedName, quorumThresholdPercentage, quorumNumbers, minNumOfOracleNetworks)
}

// RequestPriceFeedUpdate is a paid mutator transaction binding the contract method 0x9b056bc1.
//
// Solidity: function requestPriceFeedUpdate(string feedName, uint32 quorumThresholdPercentage, bytes quorumNumbers, uint8 minNumOfOracleNetworks) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) RequestPriceFeedUpdate(feedName string, quorumThresholdPercentage uint32, quorumNumbers []byte, minNumOfOracleNetworks uint8) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.RequestPriceFeedUpdate(&_ContractPriceAggregatorTaskManager.TransactOpts, feedName, quorumThresholdPercentage, quorumNumbers, minNumOfOracleNetworks)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xbafbcbe7.
//
// Solidity: function respondToTask((uint32,string,bytes,uint32,uint8) task, (uint32,uint32,uint32,string)[] taskResponses, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][])[] nonSignerStakesAndSignatures) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IPriceAggregatorTaskManagerPriceUpdateTask, taskResponses []IPriceAggregatorTaskManagerPriceUpdateTaskResponse, nonSignerStakesAndSignatures []IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "respondToTask", task, taskResponses, nonSignerStakesAndSignatures)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xbafbcbe7.
//
// Solidity: function respondToTask((uint32,string,bytes,uint32,uint8) task, (uint32,uint32,uint32,string)[] taskResponses, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][])[] nonSignerStakesAndSignatures) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) RespondToTask(task IPriceAggregatorTaskManagerPriceUpdateTask, taskResponses []IPriceAggregatorTaskManagerPriceUpdateTaskResponse, nonSignerStakesAndSignatures []IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.RespondToTask(&_ContractPriceAggregatorTaskManager.TransactOpts, task, taskResponses, nonSignerStakesAndSignatures)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xbafbcbe7.
//
// Solidity: function respondToTask((uint32,string,bytes,uint32,uint8) task, (uint32,uint32,uint32,string)[] taskResponses, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][])[] nonSignerStakesAndSignatures) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) RespondToTask(task IPriceAggregatorTaskManagerPriceUpdateTask, taskResponses []IPriceAggregatorTaskManagerPriceUpdateTaskResponse, nonSignerStakesAndSignatures []IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.RespondToTask(&_ContractPriceAggregatorTaskManager.TransactOpts, task, taskResponses, nonSignerStakesAndSignatures)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.SetPauserRegistry(&_ContractPriceAggregatorTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.SetPauserRegistry(&_ContractPriceAggregatorTaskManager.TransactOpts, newPauserRegistry)
}

// SetQuorumThreshold is a paid mutator transaction binding the contract method 0x4524c7e1.
//
// Solidity: function setQuorumThreshold(uint256 threshold) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) SetQuorumThreshold(opts *bind.TransactOpts, threshold *big.Int) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "setQuorumThreshold", threshold)
}

// SetQuorumThreshold is a paid mutator transaction binding the contract method 0x4524c7e1.
//
// Solidity: function setQuorumThreshold(uint256 threshold) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) SetQuorumThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.SetQuorumThreshold(&_ContractPriceAggregatorTaskManager.TransactOpts, threshold)
}

// SetQuorumThreshold is a paid mutator transaction binding the contract method 0x4524c7e1.
//
// Solidity: function setQuorumThreshold(uint256 threshold) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) SetQuorumThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.SetQuorumThreshold(&_ContractPriceAggregatorTaskManager.TransactOpts, threshold)
}

// SetRequiredNumberOfSources is a paid mutator transaction binding the contract method 0xcd22ed7b.
//
// Solidity: function setRequiredNumberOfSources(uint256 requiredNumOfSources) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) SetRequiredNumberOfSources(opts *bind.TransactOpts, requiredNumOfSources *big.Int) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "setRequiredNumberOfSources", requiredNumOfSources)
}

// SetRequiredNumberOfSources is a paid mutator transaction binding the contract method 0xcd22ed7b.
//
// Solidity: function setRequiredNumberOfSources(uint256 requiredNumOfSources) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) SetRequiredNumberOfSources(requiredNumOfSources *big.Int) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.SetRequiredNumberOfSources(&_ContractPriceAggregatorTaskManager.TransactOpts, requiredNumOfSources)
}

// SetRequiredNumberOfSources is a paid mutator transaction binding the contract method 0xcd22ed7b.
//
// Solidity: function setRequiredNumberOfSources(uint256 requiredNumOfSources) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) SetRequiredNumberOfSources(requiredNumOfSources *big.Int) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.SetRequiredNumberOfSources(&_ContractPriceAggregatorTaskManager.TransactOpts, requiredNumOfSources)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.SetStaleStakesForbidden(&_ContractPriceAggregatorTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.SetStaleStakesForbidden(&_ContractPriceAggregatorTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.TransferOwnership(&_ContractPriceAggregatorTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.TransferOwnership(&_ContractPriceAggregatorTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Unpause(&_ContractPriceAggregatorTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractPriceAggregatorTaskManager.Contract.Unpause(&_ContractPriceAggregatorTaskManager.TransactOpts, newPausedStatus)
}

// ContractPriceAggregatorTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerInitializedIterator struct {
	Event *ContractPriceAggregatorTaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPriceAggregatorTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorTaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPriceAggregatorTaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPriceAggregatorTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorTaskManagerInitialized represents a Initialized event raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractPriceAggregatorTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerInitializedIterator{contract: _ContractPriceAggregatorTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorTaskManagerInitialized)
				if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractPriceAggregatorTaskManagerInitialized, error) {
	event := new(ContractPriceAggregatorTaskManagerInitialized)
	if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceAggregatorTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerNewTaskCreatedIterator struct {
	Event *ContractPriceAggregatorTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPriceAggregatorTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorTaskManagerNewTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPriceAggregatorTaskManagerNewTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPriceAggregatorTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IPriceAggregatorTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c48.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256,uint32,bytes,uint32) task)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractPriceAggregatorTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerNewTaskCreatedIterator{contract: _ContractPriceAggregatorTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c48.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256,uint32,bytes,uint32) task)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorTaskManagerNewTaskCreated)
				if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTaskCreated is a log parse operation binding the contract event 0x1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c48.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256,uint32,bytes,uint32) task)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractPriceAggregatorTaskManagerNewTaskCreated, error) {
	event := new(ContractPriceAggregatorTaskManagerNewTaskCreated)
	if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceAggregatorTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerOwnershipTransferredIterator struct {
	Event *ContractPriceAggregatorTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPriceAggregatorTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorTaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPriceAggregatorTaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPriceAggregatorTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractPriceAggregatorTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerOwnershipTransferredIterator{contract: _ContractPriceAggregatorTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorTaskManagerOwnershipTransferred)
				if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractPriceAggregatorTaskManagerOwnershipTransferred, error) {
	event := new(ContractPriceAggregatorTaskManagerOwnershipTransferred)
	if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceAggregatorTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerPausedIterator struct {
	Event *ContractPriceAggregatorTaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPriceAggregatorTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorTaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPriceAggregatorTaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPriceAggregatorTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorTaskManagerPaused represents a Paused event raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractPriceAggregatorTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerPausedIterator{contract: _ContractPriceAggregatorTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorTaskManagerPaused)
				if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) ParsePaused(log types.Log) (*ContractPriceAggregatorTaskManagerPaused, error) {
	event := new(ContractPriceAggregatorTaskManagerPaused)
	if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceAggregatorTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerPauserRegistrySetIterator struct {
	Event *ContractPriceAggregatorTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPriceAggregatorTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorTaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPriceAggregatorTaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPriceAggregatorTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractPriceAggregatorTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerPauserRegistrySetIterator{contract: _ContractPriceAggregatorTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorTaskManagerPauserRegistrySet)
				if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractPriceAggregatorTaskManagerPauserRegistrySet, error) {
	event := new(ContractPriceAggregatorTaskManagerPauserRegistrySet)
	if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceAggregatorTaskManagerPriceUpdateRequestedIterator is returned from FilterPriceUpdateRequested and is used to iterate over the raw logs and unpacked data for PriceUpdateRequested events raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerPriceUpdateRequestedIterator struct {
	Event *ContractPriceAggregatorTaskManagerPriceUpdateRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPriceAggregatorTaskManagerPriceUpdateRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorTaskManagerPriceUpdateRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPriceAggregatorTaskManagerPriceUpdateRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPriceAggregatorTaskManagerPriceUpdateRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorTaskManagerPriceUpdateRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorTaskManagerPriceUpdateRequested represents a PriceUpdateRequested event raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerPriceUpdateRequested struct {
	TaskIndex uint32
	Task      IPriceAggregatorTaskManagerPriceUpdateTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdateRequested is a free log retrieval operation binding the contract event 0x903a8215cb7ce14c9230c13623ee3c35bc42c4a0cb4fed2ff3a24ca32c7773f3.
//
// Solidity: event PriceUpdateRequested(uint32 indexed taskIndex, (uint32,string,bytes,uint32,uint8) task)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) FilterPriceUpdateRequested(opts *bind.FilterOpts, taskIndex []uint32) (*ContractPriceAggregatorTaskManagerPriceUpdateRequestedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.FilterLogs(opts, "PriceUpdateRequested", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerPriceUpdateRequestedIterator{contract: _ContractPriceAggregatorTaskManager.contract, event: "PriceUpdateRequested", logs: logs, sub: sub}, nil
}

// WatchPriceUpdateRequested is a free log subscription operation binding the contract event 0x903a8215cb7ce14c9230c13623ee3c35bc42c4a0cb4fed2ff3a24ca32c7773f3.
//
// Solidity: event PriceUpdateRequested(uint32 indexed taskIndex, (uint32,string,bytes,uint32,uint8) task)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) WatchPriceUpdateRequested(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorTaskManagerPriceUpdateRequested, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.WatchLogs(opts, "PriceUpdateRequested", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorTaskManagerPriceUpdateRequested)
				if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "PriceUpdateRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceUpdateRequested is a log parse operation binding the contract event 0x903a8215cb7ce14c9230c13623ee3c35bc42c4a0cb4fed2ff3a24ca32c7773f3.
//
// Solidity: event PriceUpdateRequested(uint32 indexed taskIndex, (uint32,string,bytes,uint32,uint8) task)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) ParsePriceUpdateRequested(log types.Log) (*ContractPriceAggregatorTaskManagerPriceUpdateRequested, error) {
	event := new(ContractPriceAggregatorTaskManagerPriceUpdateRequested)
	if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "PriceUpdateRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractPriceAggregatorTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractPriceAggregatorTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceAggregatorTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractPriceAggregatorTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPriceAggregatorTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorTaskManagerTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPriceAggregatorTaskManagerTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPriceAggregatorTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractPriceAggregatorTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractPriceAggregatorTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorTaskManagerTaskChallengedSuccessfully)
				if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractPriceAggregatorTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractPriceAggregatorTaskManagerTaskChallengedSuccessfully)
	if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractPriceAggregatorTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractPriceAggregatorTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceAggregatorTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerTaskCompletedIterator struct {
	Event *ContractPriceAggregatorTaskManagerTaskCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPriceAggregatorTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorTaskManagerTaskCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPriceAggregatorTaskManagerTaskCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPriceAggregatorTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractPriceAggregatorTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerTaskCompletedIterator{contract: _ContractPriceAggregatorTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorTaskManagerTaskCompleted)
				if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractPriceAggregatorTaskManagerTaskCompleted, error) {
	event := new(ContractPriceAggregatorTaskManagerTaskCompleted)
	if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceAggregatorTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerTaskRespondedIterator struct {
	Event *ContractPriceAggregatorTaskManagerTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPriceAggregatorTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorTaskManagerTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPriceAggregatorTaskManagerTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPriceAggregatorTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorTaskManagerTaskResponded represents a TaskResponded event raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerTaskResponded struct {
	TaskIndex            uint32
	TaskResponse         IPriceAggregatorTaskManagerPriceUpdateTaskResponse
	TaskResponseMetadata IPriceAggregatorTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xbf877b433f8d349e650f5f2d05288ff1b727ac7652853301d6c73558b2757f7f.
//
// Solidity: event TaskResponded(uint32 indexed taskIndex, (uint32,uint32,uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts, taskIndex []uint32) (*ContractPriceAggregatorTaskManagerTaskRespondedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.FilterLogs(opts, "TaskResponded", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerTaskRespondedIterator{contract: _ContractPriceAggregatorTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xbf877b433f8d349e650f5f2d05288ff1b727ac7652853301d6c73558b2757f7f.
//
// Solidity: event TaskResponded(uint32 indexed taskIndex, (uint32,uint32,uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorTaskManagerTaskResponded, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.WatchLogs(opts, "TaskResponded", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorTaskManagerTaskResponded)
				if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskResponded is a log parse operation binding the contract event 0xbf877b433f8d349e650f5f2d05288ff1b727ac7652853301d6c73558b2757f7f.
//
// Solidity: event TaskResponded(uint32 indexed taskIndex, (uint32,uint32,uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractPriceAggregatorTaskManagerTaskResponded, error) {
	event := new(ContractPriceAggregatorTaskManagerTaskResponded)
	if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceAggregatorTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerUnpausedIterator struct {
	Event *ContractPriceAggregatorTaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPriceAggregatorTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorTaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPriceAggregatorTaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPriceAggregatorTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorTaskManagerUnpaused represents a Unpaused event raised by the ContractPriceAggregatorTaskManager contract.
type ContractPriceAggregatorTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractPriceAggregatorTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorTaskManagerUnpausedIterator{contract: _ContractPriceAggregatorTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractPriceAggregatorTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorTaskManagerUnpaused)
				if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractPriceAggregatorTaskManager *ContractPriceAggregatorTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractPriceAggregatorTaskManagerUnpaused, error) {
	event := new(ContractPriceAggregatorTaskManagerUnpaused)
	if err := _ContractPriceAggregatorTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
