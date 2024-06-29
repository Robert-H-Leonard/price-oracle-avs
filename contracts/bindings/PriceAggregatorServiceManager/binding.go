// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractPriceAggregatorServiceManager

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

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	StrategiesAndMultipliers []IPaymentCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IPaymentCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractPriceAggregatorServiceManagerMetaData contains all meta data concerning the ContractPriceAggregatorServiceManager contract.
var ContractPriceAggregatorServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_PriceAggregatorTaskManager\",\"type\":\"address\",\"internalType\":\"contractIPriceAggregatorTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"PriceAggregatorTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPriceAggregatorTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fetchOperatorUrls\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"httpUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rpcUrl\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidOperator\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorHttpUrls\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorRpcUrls\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorConsensusUrl\",\"inputs\":[{\"name\":\"httpUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rpcUrl\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUrlRegistered\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"httpUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b50604051620023b9380380620023b983398101604081905262000035916200015a565b6001600160a01b03808516608052600060a081905281851660c05290831660e05284908484620000646200007f565b505050506001600160a01b03166101005250620001c2915050565b600054610100900460ff1615620000ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200013f576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015757600080fd5b50565b600080600080608085870312156200017157600080fd5b84516200017e8162000141565b6020860151909450620001918162000141565b6040860151909350620001a48162000141565b6060860151909250620001b78162000141565b939692955090935050565b60805160a05160c05160e0516101005161213462000285600039600081816101ee0152610a1c01526000818161073b015281816108970152818161092e015281816112940152818161141801526114b7015260008181610566015281816105f50152818161067501528181610adf01528181610cb301528181610d4901528181611158015281816111d201526113730152600081816103f601526104d401526000818161019601528181610d0701528181610da501526110ff01526121346000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063825d6f7d116100a2578063a803be5711610071578063a803be5714610247578063a98fb35514610268578063e2d53d7c1461027b578063e481af9d1461029e578063f2fde38b146102a657600080fd5b8063825d6f7d146101e95780638da5cb5b146102105780639926ee7d14610221578063a364f4da1461023457600080fd5b806360707ef6116100de57806360707ef6146101745780636b3aa72e146101945780636fcb7a75146101ce578063715018a6146101e157600080fd5b80631b4455161461011057806333cfb7b71461012557806338c8ee641461014e5780634fa7ab3614610161575b600080fd5b61012361011e366004611845565b6102b9565b005b6101386101333660046118cf565b610541565b60405161014591906118f3565b60405180910390f35b61012361015c3660046118cf565b610a11565b61012361016f366004611a38565b610ac7565b6101876101823660046118cf565b610be1565b6040516101459190611ae9565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b039091168152602001610145565b6101876101dc3660046118cf565b610c7b565b610123610c94565b6101b67f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b03166101b6565b61012361022f366004611afc565b610ca8565b6101236102423660046118cf565b610d3e565b61025a6102553660046118cf565b610e05565b604051610145929190611ba7565b610123610276366004611bd5565b6110e0565b61028e6102893660046118cf565b611134565b6040519015158152602001610145565b6101386111cc565b6101236102b43660046118cf565b611596565b6102c161160c565b60005b818110156104bc578282828181106102de576102de611c12565b90506020028101906102f09190611c28565b6103019060408101906020016118cf565b6001600160a01b03166323b872dd333086868681811061032357610323611c12565b90506020028101906103359190611c28565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af115801561038c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b09190611c58565b508282828181106103c3576103c3611c12565b90506020028101906103d59190611c28565b6103e69060408101906020016118cf565b6001600160a01b031663095ea7b37f000000000000000000000000000000000000000000000000000000000000000085858581811061042757610427611c12565b90506020028101906104399190611c28565b604080516001600160e01b031960e086901b1681526001600160a01b039093166004840152013560248201526044016020604051808303816000875af1158015610487573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ab9190611c58565b506104b581611c90565b90506102c4565b50604051630da22a8b60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631b4455169061050b9085908590600401611d44565b600060405180830381600087803b15801561052557600080fd5b505af1158015610539573d6000803e3d6000fd5b505050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156105ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105d19190611e52565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561063c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106609190611e6b565b90506001600160c01b03811615806106fa57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f59190611e94565b60ff16155b1561071657505060408051600081526020810190915292915050565b600061072a826001600160c01b0316611666565b90506000805b8251811015610800577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f584838151811061077a5761077a611c12565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156107be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e29190611e52565b6107ec9083611eb7565b9150806107f881611c90565b915050610730565b5060008167ffffffffffffffff81111561081c5761081c611940565b604051908082528060200260200182016040528015610845578160200160208202803683370190505b5090506000805b8451811015610a0457600085828151811061086957610869611c12565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa1580156108de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109029190611e52565b905060005b818110156109ee576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561097c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a09190611ecf565b600001518686815181106109b6576109b6611c12565b6001600160a01b0390921660209283029190910190910152846109d881611c90565b95505080806109e690611c90565b915050610907565b50505080806109fc90611c90565b91505061084c565b5090979650505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ac45760405162461bcd60e51b815260206004820152604760248201527f6f6e6c79507269636541676772656761746f725461736b4d616e616765723a2060448201527f6e6f742066726f6d206372656469626c65207371756172696e67207461736b2060648201526636b0b730b3b2b960c91b608482015260a4015b60405180910390fd5b50565b604051631619718360e21b81523360048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635865c60c906024016040805180830381865afa158015610b2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b519190611f10565b8051909150610b5f57600080fd5b3360009081526097602090815260409091208451610b7f928601906117ac565b503360009081526098602090815260409091208351610ba0928501906117ac565b507f100388966d784f0d0a7df3e4926797840162cc1e2cf828c6ca6ec6adaeb4640f338484604051610bd493929190611f40565b60405180910390a1505050565b60976020526000908152604090208054610bfa90611f80565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2690611f80565b8015610c735780601f10610c4857610100808354040283529160200191610c73565b820191906000526020600020905b815481529060010190602001808311610c5657829003601f168201915b505050505081565b60986020526000908152604090208054610bfa90611f80565b610c9c61160c565b610ca66000611729565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610cf05760405162461bcd60e51b8152600401610abb90611fbb565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061050b9085908590600401612033565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610d865760405162461bcd60e51b8152600401610abb90611fbb565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610dea57600080fd5b505af1158015610dfe573d6000803e3d6000fd5b5050505050565b6001600160a01b0381166000908152609760205260408120805460609283929091610e2f90611f80565b80601f0160208091040260200160405190810160405280929190818152602001828054610e5b90611f80565b8015610ea85780601f10610e7d57610100808354040283529160200191610ea8565b820191906000526020600020905b815481529060010190602001808311610e8b57829003601f168201915b5050506001600160a01b038716600090815260986020526040812080549495509093909250610ed79150611f80565b80601f0160208091040260200160405190810160405280929190818152602001828054610f0390611f80565b8015610f505780601f10610f2557610100808354040283529160200191610f50565b820191906000526020600020905b815481529060010190602001808311610f3357829003601f168201915b505050505090506000825111610f785760405162461bcd60e51b8152600401610abb9061207e565b6000815111610f995760405162461bcd60e51b8152600401610abb9061207e565b6001600160a01b0385166000908152609760209081526040808320609890925290912081548290610fc990611f80565b80601f0160208091040260200160405190810160405280929190818152602001828054610ff590611f80565b80156110425780601f1061101757610100808354040283529160200191611042565b820191906000526020600020905b81548152906001019060200180831161102557829003601f168201915b5050505050915080805461105590611f80565b80601f016020809104026020016040519081016040528092919081815260200182805461108190611f80565b80156110ce5780601f106110a3576101008083540402835291602001916110ce565b820191906000526020600020905b8154815290600101906020018083116110b157829003601f168201915b50505050509050935093505050915091565b6110e861160c565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610dd0908490600401611ae9565b604051631619718360e21b81526001600160a01b03828116600483015260009182917f00000000000000000000000000000000000000000000000000000000000000001690635865c60c906024016040805180830381865afa15801561119e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111c29190611f10565b5115159392505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561122e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112529190611e94565b60ff1690508061127057505060408051600081526020810190915290565b6000805b8281101561132557604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa1580156112e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113079190611e52565b6113119083611eb7565b91508061131d81611c90565b915050611274565b5060008167ffffffffffffffff81111561134157611341611940565b60405190808252806020026020018201604052801561136a578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113f39190611e94565b60ff1681101561158c57604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015611467573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061148b9190611e52565b905060005b81811015611577576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015611505573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115299190611ecf565b6000015185858151811061153f5761153f611c12565b6001600160a01b03909216602092830291909101909101528361156181611c90565b945050808061156f90611c90565b915050611490565b5050808061158490611c90565b915050611371565b5090949350505050565b61159e61160c565b6001600160a01b0381166116035760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610abb565b610ac481611729565b6033546001600160a01b03163314610ca65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610abb565b60606000806116748461177b565b61ffff1667ffffffffffffffff81111561169057611690611940565b6040519080825280601f01601f1916602001820160405280156116ba576020820181803683370190505b5090506000805b8251821080156116d2575061010081105b1561158c576001811b935085841615611719578060f81b8383815181106116fb576116fb611c12565b60200101906001600160f81b031916908160001a9053508160010191505b61172281611c90565b90506116c1565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b82156117a6576117906001846120c5565b909216918061179e816120dc565b91505061177f565b92915050565b8280546117b890611f80565b90600052602060002090601f0160209004810192826117da5760008555611820565b82601f106117f357805160ff1916838001178555611820565b82800160010185558215611820579182015b82811115611820578251825591602001919060010190611805565b5061182c929150611830565b5090565b5b8082111561182c5760008155600101611831565b6000806020838503121561185857600080fd5b823567ffffffffffffffff8082111561187057600080fd5b818501915085601f83011261188457600080fd5b81358181111561189357600080fd5b8660208260051b85010111156118a857600080fd5b60209290920196919550909350505050565b6001600160a01b0381168114610ac457600080fd5b6000602082840312156118e157600080fd5b81356118ec816118ba565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156119345783516001600160a01b03168352928401929184019160010161190f565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561197957611979611940565b60405290565b6040805190810167ffffffffffffffff8111828210171561197957611979611940565b600067ffffffffffffffff808411156119bd576119bd611940565b604051601f8501601f19908116603f011681019082821181831017156119e5576119e5611940565b816040528093508581528686860111156119fe57600080fd5b858560208301376000602087830101525050509392505050565b600082601f830112611a2957600080fd5b6118ec838335602085016119a2565b60008060408385031215611a4b57600080fd5b823567ffffffffffffffff80821115611a6357600080fd5b611a6f86838701611a18565b93506020850135915080821115611a8557600080fd5b50611a9285828601611a18565b9150509250929050565b6000815180845260005b81811015611ac257602081850181015186830182015201611aa6565b81811115611ad4576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006118ec6020830184611a9c565b60008060408385031215611b0f57600080fd5b8235611b1a816118ba565b9150602083013567ffffffffffffffff80821115611b3757600080fd5b9084019060608287031215611b4b57600080fd5b611b53611956565b823582811115611b6257600080fd5b83019150601f82018713611b7557600080fd5b611b84878335602085016119a2565b815260208301356020820152604083013560408201528093505050509250929050565b604081526000611bba6040830185611a9c565b8281036020840152611bcc8185611a9c565b95945050505050565b600060208284031215611be757600080fd5b813567ffffffffffffffff811115611bfe57600080fd5b611c0a84828501611a18565b949350505050565b634e487b7160e01b600052603260045260246000fd5b60008235609e19833603018112611c3e57600080fd5b9190910192915050565b8035611c53816118ba565b919050565b600060208284031215611c6a57600080fd5b815180151581146118ec57600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611ca457611ca4611c7a565b5060010190565b6bffffffffffffffffffffffff81168114610ac457600080fd5b8183526000602080850194508260005b85811015611d25578135611ce8816118ba565b6001600160a01b0316875281830135611d0081611cab565b6bffffffffffffffffffffffff16878401526040968701969190910190600101611cd5565b509495945050505050565b803563ffffffff81168114611c5357600080fd5b60208082528181018390526000906040808401600586901b8501820187855b88811015611e4457878303603f190184528135368b9003609e19018112611d8957600080fd5b8a0160a0813536839003601e19018112611da257600080fd5b8201803567ffffffffffffffff811115611dbb57600080fd5b8060061b3603841315611dcd57600080fd5b828752611ddf838801828c8501611cc5565b92505050611dee888301611c48565b6001600160a01b03168886015281870135878601526060611e10818401611d30565b63ffffffff16908601526080611e27838201611d30565b63ffffffff16950194909452509285019290850190600101611d63565b509098975050505050505050565b600060208284031215611e6457600080fd5b5051919050565b600060208284031215611e7d57600080fd5b81516001600160c01b03811681146118ec57600080fd5b600060208284031215611ea657600080fd5b815160ff811681146118ec57600080fd5b60008219821115611eca57611eca611c7a565b500190565b600060408284031215611ee157600080fd5b611ee961197f565b8251611ef4816118ba565b81526020830151611f0481611cab565b60208201529392505050565b600060408284031215611f2257600080fd5b611f2a61197f565b82518152602083015160038110611f0457600080fd5b6001600160a01b0384168152606060208201819052600090611f6490830185611a9c565b8281036040840152611f768185611a9c565b9695505050505050565b600181811c90821680611f9457607f821691505b60208210811415611fb557634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b038316815260406020820152600082516060604084015261205d60a0840182611a9c565b90506020840151606084015260408401516080840152809150509392505050565b60208082526027908201527f4e6f2075726c207265676973746572656420666f7220726571756573746564206040820152666164647265737360c81b606082015260800190565b6000828210156120d7576120d7611c7a565b500390565b600061ffff808316818114156120f4576120f4611c7a565b600101939250505056fea26469706673582212201f2cfe6ba5a3f29de5bf08d3d3b4178c8924cf65b51b53eae0ea143c98925b8164736f6c634300080c0033",
}

// ContractPriceAggregatorServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractPriceAggregatorServiceManagerMetaData.ABI instead.
var ContractPriceAggregatorServiceManagerABI = ContractPriceAggregatorServiceManagerMetaData.ABI

// ContractPriceAggregatorServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractPriceAggregatorServiceManagerMetaData.Bin instead.
var ContractPriceAggregatorServiceManagerBin = ContractPriceAggregatorServiceManagerMetaData.Bin

// DeployContractPriceAggregatorServiceManager deploys a new Ethereum contract, binding an instance of ContractPriceAggregatorServiceManager to it.
func DeployContractPriceAggregatorServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _PriceAggregatorTaskManager common.Address) (common.Address, *types.Transaction, *ContractPriceAggregatorServiceManager, error) {
	parsed, err := ContractPriceAggregatorServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractPriceAggregatorServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _PriceAggregatorTaskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractPriceAggregatorServiceManager{ContractPriceAggregatorServiceManagerCaller: ContractPriceAggregatorServiceManagerCaller{contract: contract}, ContractPriceAggregatorServiceManagerTransactor: ContractPriceAggregatorServiceManagerTransactor{contract: contract}, ContractPriceAggregatorServiceManagerFilterer: ContractPriceAggregatorServiceManagerFilterer{contract: contract}}, nil
}

// ContractPriceAggregatorServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractPriceAggregatorServiceManager struct {
	ContractPriceAggregatorServiceManagerCaller     // Read-only binding to the contract
	ContractPriceAggregatorServiceManagerTransactor // Write-only binding to the contract
	ContractPriceAggregatorServiceManagerFilterer   // Log filterer for contract events
}

// ContractPriceAggregatorServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractPriceAggregatorServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPriceAggregatorServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractPriceAggregatorServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPriceAggregatorServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractPriceAggregatorServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPriceAggregatorServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractPriceAggregatorServiceManagerSession struct {
	Contract     *ContractPriceAggregatorServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                          // Call options to use throughout this session
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractPriceAggregatorServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractPriceAggregatorServiceManagerCallerSession struct {
	Contract *ContractPriceAggregatorServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                // Call options to use throughout this session
}

// ContractPriceAggregatorServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractPriceAggregatorServiceManagerTransactorSession struct {
	Contract     *ContractPriceAggregatorServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                // Transaction auth options to use throughout this session
}

// ContractPriceAggregatorServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractPriceAggregatorServiceManagerRaw struct {
	Contract *ContractPriceAggregatorServiceManager // Generic contract binding to access the raw methods on
}

// ContractPriceAggregatorServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractPriceAggregatorServiceManagerCallerRaw struct {
	Contract *ContractPriceAggregatorServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractPriceAggregatorServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractPriceAggregatorServiceManagerTransactorRaw struct {
	Contract *ContractPriceAggregatorServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractPriceAggregatorServiceManager creates a new instance of ContractPriceAggregatorServiceManager, bound to a specific deployed contract.
func NewContractPriceAggregatorServiceManager(address common.Address, backend bind.ContractBackend) (*ContractPriceAggregatorServiceManager, error) {
	contract, err := bindContractPriceAggregatorServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorServiceManager{ContractPriceAggregatorServiceManagerCaller: ContractPriceAggregatorServiceManagerCaller{contract: contract}, ContractPriceAggregatorServiceManagerTransactor: ContractPriceAggregatorServiceManagerTransactor{contract: contract}, ContractPriceAggregatorServiceManagerFilterer: ContractPriceAggregatorServiceManagerFilterer{contract: contract}}, nil
}

// NewContractPriceAggregatorServiceManagerCaller creates a new read-only instance of ContractPriceAggregatorServiceManager, bound to a specific deployed contract.
func NewContractPriceAggregatorServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractPriceAggregatorServiceManagerCaller, error) {
	contract, err := bindContractPriceAggregatorServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorServiceManagerCaller{contract: contract}, nil
}

// NewContractPriceAggregatorServiceManagerTransactor creates a new write-only instance of ContractPriceAggregatorServiceManager, bound to a specific deployed contract.
func NewContractPriceAggregatorServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractPriceAggregatorServiceManagerTransactor, error) {
	contract, err := bindContractPriceAggregatorServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorServiceManagerTransactor{contract: contract}, nil
}

// NewContractPriceAggregatorServiceManagerFilterer creates a new log filterer instance of ContractPriceAggregatorServiceManager, bound to a specific deployed contract.
func NewContractPriceAggregatorServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractPriceAggregatorServiceManagerFilterer, error) {
	contract, err := bindContractPriceAggregatorServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorServiceManagerFilterer{contract: contract}, nil
}

// bindContractPriceAggregatorServiceManager binds a generic wrapper to an already deployed contract.
func bindContractPriceAggregatorServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractPriceAggregatorServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractPriceAggregatorServiceManager.Contract.ContractPriceAggregatorServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.ContractPriceAggregatorServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.ContractPriceAggregatorServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractPriceAggregatorServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.contract.Transact(opts, method, params...)
}

// PriceAggregatorTaskManager is a free data retrieval call binding the contract method 0x825d6f7d.
//
// Solidity: function PriceAggregatorTaskManager() view returns(address)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCaller) PriceAggregatorTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceAggregatorServiceManager.contract.Call(opts, &out, "PriceAggregatorTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceAggregatorTaskManager is a free data retrieval call binding the contract method 0x825d6f7d.
//
// Solidity: function PriceAggregatorTaskManager() view returns(address)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) PriceAggregatorTaskManager() (common.Address, error) {
	return _ContractPriceAggregatorServiceManager.Contract.PriceAggregatorTaskManager(&_ContractPriceAggregatorServiceManager.CallOpts)
}

// PriceAggregatorTaskManager is a free data retrieval call binding the contract method 0x825d6f7d.
//
// Solidity: function PriceAggregatorTaskManager() view returns(address)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCallerSession) PriceAggregatorTaskManager() (common.Address, error) {
	return _ContractPriceAggregatorServiceManager.Contract.PriceAggregatorTaskManager(&_ContractPriceAggregatorServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceAggregatorServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractPriceAggregatorServiceManager.Contract.AvsDirectory(&_ContractPriceAggregatorServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractPriceAggregatorServiceManager.Contract.AvsDirectory(&_ContractPriceAggregatorServiceManager.CallOpts)
}

// FetchOperatorUrls is a free data retrieval call binding the contract method 0xa803be57.
//
// Solidity: function fetchOperatorUrls(address operatorAddress) view returns(string httpUrl, string rpcUrl)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCaller) FetchOperatorUrls(opts *bind.CallOpts, operatorAddress common.Address) (struct {
	HttpUrl string
	RpcUrl  string
}, error) {
	var out []interface{}
	err := _ContractPriceAggregatorServiceManager.contract.Call(opts, &out, "fetchOperatorUrls", operatorAddress)

	outstruct := new(struct {
		HttpUrl string
		RpcUrl  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HttpUrl = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.RpcUrl = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// FetchOperatorUrls is a free data retrieval call binding the contract method 0xa803be57.
//
// Solidity: function fetchOperatorUrls(address operatorAddress) view returns(string httpUrl, string rpcUrl)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) FetchOperatorUrls(operatorAddress common.Address) (struct {
	HttpUrl string
	RpcUrl  string
}, error) {
	return _ContractPriceAggregatorServiceManager.Contract.FetchOperatorUrls(&_ContractPriceAggregatorServiceManager.CallOpts, operatorAddress)
}

// FetchOperatorUrls is a free data retrieval call binding the contract method 0xa803be57.
//
// Solidity: function fetchOperatorUrls(address operatorAddress) view returns(string httpUrl, string rpcUrl)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCallerSession) FetchOperatorUrls(operatorAddress common.Address) (struct {
	HttpUrl string
	RpcUrl  string
}, error) {
	return _ContractPriceAggregatorServiceManager.Contract.FetchOperatorUrls(&_ContractPriceAggregatorServiceManager.CallOpts, operatorAddress)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractPriceAggregatorServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractPriceAggregatorServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractPriceAggregatorServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractPriceAggregatorServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractPriceAggregatorServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractPriceAggregatorServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractPriceAggregatorServiceManager.Contract.GetRestakeableStrategies(&_ContractPriceAggregatorServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractPriceAggregatorServiceManager.Contract.GetRestakeableStrategies(&_ContractPriceAggregatorServiceManager.CallOpts)
}

// IsValidOperator is a free data retrieval call binding the contract method 0xe2d53d7c.
//
// Solidity: function isValidOperator(address operatorAddress) view returns(bool)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCaller) IsValidOperator(opts *bind.CallOpts, operatorAddress common.Address) (bool, error) {
	var out []interface{}
	err := _ContractPriceAggregatorServiceManager.contract.Call(opts, &out, "isValidOperator", operatorAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidOperator is a free data retrieval call binding the contract method 0xe2d53d7c.
//
// Solidity: function isValidOperator(address operatorAddress) view returns(bool)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) IsValidOperator(operatorAddress common.Address) (bool, error) {
	return _ContractPriceAggregatorServiceManager.Contract.IsValidOperator(&_ContractPriceAggregatorServiceManager.CallOpts, operatorAddress)
}

// IsValidOperator is a free data retrieval call binding the contract method 0xe2d53d7c.
//
// Solidity: function isValidOperator(address operatorAddress) view returns(bool)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCallerSession) IsValidOperator(operatorAddress common.Address) (bool, error) {
	return _ContractPriceAggregatorServiceManager.Contract.IsValidOperator(&_ContractPriceAggregatorServiceManager.CallOpts, operatorAddress)
}

// OperatorHttpUrls is a free data retrieval call binding the contract method 0x60707ef6.
//
// Solidity: function operatorHttpUrls(address ) view returns(string)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCaller) OperatorHttpUrls(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _ContractPriceAggregatorServiceManager.contract.Call(opts, &out, "operatorHttpUrls", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OperatorHttpUrls is a free data retrieval call binding the contract method 0x60707ef6.
//
// Solidity: function operatorHttpUrls(address ) view returns(string)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) OperatorHttpUrls(arg0 common.Address) (string, error) {
	return _ContractPriceAggregatorServiceManager.Contract.OperatorHttpUrls(&_ContractPriceAggregatorServiceManager.CallOpts, arg0)
}

// OperatorHttpUrls is a free data retrieval call binding the contract method 0x60707ef6.
//
// Solidity: function operatorHttpUrls(address ) view returns(string)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCallerSession) OperatorHttpUrls(arg0 common.Address) (string, error) {
	return _ContractPriceAggregatorServiceManager.Contract.OperatorHttpUrls(&_ContractPriceAggregatorServiceManager.CallOpts, arg0)
}

// OperatorRpcUrls is a free data retrieval call binding the contract method 0x6fcb7a75.
//
// Solidity: function operatorRpcUrls(address ) view returns(string)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCaller) OperatorRpcUrls(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _ContractPriceAggregatorServiceManager.contract.Call(opts, &out, "operatorRpcUrls", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OperatorRpcUrls is a free data retrieval call binding the contract method 0x6fcb7a75.
//
// Solidity: function operatorRpcUrls(address ) view returns(string)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) OperatorRpcUrls(arg0 common.Address) (string, error) {
	return _ContractPriceAggregatorServiceManager.Contract.OperatorRpcUrls(&_ContractPriceAggregatorServiceManager.CallOpts, arg0)
}

// OperatorRpcUrls is a free data retrieval call binding the contract method 0x6fcb7a75.
//
// Solidity: function operatorRpcUrls(address ) view returns(string)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCallerSession) OperatorRpcUrls(arg0 common.Address) (string, error) {
	return _ContractPriceAggregatorServiceManager.Contract.OperatorRpcUrls(&_ContractPriceAggregatorServiceManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceAggregatorServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) Owner() (common.Address, error) {
	return _ContractPriceAggregatorServiceManager.Contract.Owner(&_ContractPriceAggregatorServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractPriceAggregatorServiceManager.Contract.Owner(&_ContractPriceAggregatorServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractPriceAggregatorServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractPriceAggregatorServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.FreezeOperator(&_ContractPriceAggregatorServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.FreezeOperator(&_ContractPriceAggregatorServiceManager.TransactOpts, operatorAddr)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.PayForRange(&_ContractPriceAggregatorServiceManager.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.PayForRange(&_ContractPriceAggregatorServiceManager.TransactOpts, rangePayments)
}

// RegisterOperatorConsensusUrl is a paid mutator transaction binding the contract method 0x4fa7ab36.
//
// Solidity: function registerOperatorConsensusUrl(string httpUrl, string rpcUrl) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactor) RegisterOperatorConsensusUrl(opts *bind.TransactOpts, httpUrl string, rpcUrl string) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.contract.Transact(opts, "registerOperatorConsensusUrl", httpUrl, rpcUrl)
}

// RegisterOperatorConsensusUrl is a paid mutator transaction binding the contract method 0x4fa7ab36.
//
// Solidity: function registerOperatorConsensusUrl(string httpUrl, string rpcUrl) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) RegisterOperatorConsensusUrl(httpUrl string, rpcUrl string) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.RegisterOperatorConsensusUrl(&_ContractPriceAggregatorServiceManager.TransactOpts, httpUrl, rpcUrl)
}

// RegisterOperatorConsensusUrl is a paid mutator transaction binding the contract method 0x4fa7ab36.
//
// Solidity: function registerOperatorConsensusUrl(string httpUrl, string rpcUrl) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactorSession) RegisterOperatorConsensusUrl(httpUrl string, rpcUrl string) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.RegisterOperatorConsensusUrl(&_ContractPriceAggregatorServiceManager.TransactOpts, httpUrl, rpcUrl)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.RegisterOperatorToAVS(&_ContractPriceAggregatorServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.RegisterOperatorToAVS(&_ContractPriceAggregatorServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.RenounceOwnership(&_ContractPriceAggregatorServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.RenounceOwnership(&_ContractPriceAggregatorServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.TransferOwnership(&_ContractPriceAggregatorServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.TransferOwnership(&_ContractPriceAggregatorServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.UpdateAVSMetadataURI(&_ContractPriceAggregatorServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractPriceAggregatorServiceManager.Contract.UpdateAVSMetadataURI(&_ContractPriceAggregatorServiceManager.TransactOpts, _metadataURI)
}

// ContractPriceAggregatorServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractPriceAggregatorServiceManager contract.
type ContractPriceAggregatorServiceManagerInitializedIterator struct {
	Event *ContractPriceAggregatorServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractPriceAggregatorServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorServiceManagerInitialized)
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
		it.Event = new(ContractPriceAggregatorServiceManagerInitialized)
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
func (it *ContractPriceAggregatorServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorServiceManagerInitialized represents a Initialized event raised by the ContractPriceAggregatorServiceManager contract.
type ContractPriceAggregatorServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractPriceAggregatorServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractPriceAggregatorServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorServiceManagerInitializedIterator{contract: _ContractPriceAggregatorServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractPriceAggregatorServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorServiceManagerInitialized)
				if err := _ContractPriceAggregatorServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractPriceAggregatorServiceManagerInitialized, error) {
	event := new(ContractPriceAggregatorServiceManagerInitialized)
	if err := _ContractPriceAggregatorServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceAggregatorServiceManagerOperatorUrlRegisteredIterator is returned from FilterOperatorUrlRegistered and is used to iterate over the raw logs and unpacked data for OperatorUrlRegistered events raised by the ContractPriceAggregatorServiceManager contract.
type ContractPriceAggregatorServiceManagerOperatorUrlRegisteredIterator struct {
	Event *ContractPriceAggregatorServiceManagerOperatorUrlRegistered // Event containing the contract specifics and raw log

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
func (it *ContractPriceAggregatorServiceManagerOperatorUrlRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorServiceManagerOperatorUrlRegistered)
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
		it.Event = new(ContractPriceAggregatorServiceManagerOperatorUrlRegistered)
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
func (it *ContractPriceAggregatorServiceManagerOperatorUrlRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorServiceManagerOperatorUrlRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorServiceManagerOperatorUrlRegistered represents a OperatorUrlRegistered event raised by the ContractPriceAggregatorServiceManager contract.
type ContractPriceAggregatorServiceManagerOperatorUrlRegistered struct {
	OperatorId common.Address
	HttpUrl    string
	RpcUrl     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorUrlRegistered is a free log retrieval operation binding the contract event 0x100388966d784f0d0a7df3e4926797840162cc1e2cf828c6ca6ec6adaeb4640f.
//
// Solidity: event OperatorUrlRegistered(address operatorId, string httpUrl, string rpcUrl)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerFilterer) FilterOperatorUrlRegistered(opts *bind.FilterOpts) (*ContractPriceAggregatorServiceManagerOperatorUrlRegisteredIterator, error) {

	logs, sub, err := _ContractPriceAggregatorServiceManager.contract.FilterLogs(opts, "OperatorUrlRegistered")
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorServiceManagerOperatorUrlRegisteredIterator{contract: _ContractPriceAggregatorServiceManager.contract, event: "OperatorUrlRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorUrlRegistered is a free log subscription operation binding the contract event 0x100388966d784f0d0a7df3e4926797840162cc1e2cf828c6ca6ec6adaeb4640f.
//
// Solidity: event OperatorUrlRegistered(address operatorId, string httpUrl, string rpcUrl)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerFilterer) WatchOperatorUrlRegistered(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorServiceManagerOperatorUrlRegistered) (event.Subscription, error) {

	logs, sub, err := _ContractPriceAggregatorServiceManager.contract.WatchLogs(opts, "OperatorUrlRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorServiceManagerOperatorUrlRegistered)
				if err := _ContractPriceAggregatorServiceManager.contract.UnpackLog(event, "OperatorUrlRegistered", log); err != nil {
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

// ParseOperatorUrlRegistered is a log parse operation binding the contract event 0x100388966d784f0d0a7df3e4926797840162cc1e2cf828c6ca6ec6adaeb4640f.
//
// Solidity: event OperatorUrlRegistered(address operatorId, string httpUrl, string rpcUrl)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerFilterer) ParseOperatorUrlRegistered(log types.Log) (*ContractPriceAggregatorServiceManagerOperatorUrlRegistered, error) {
	event := new(ContractPriceAggregatorServiceManagerOperatorUrlRegistered)
	if err := _ContractPriceAggregatorServiceManager.contract.UnpackLog(event, "OperatorUrlRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceAggregatorServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractPriceAggregatorServiceManager contract.
type ContractPriceAggregatorServiceManagerOwnershipTransferredIterator struct {
	Event *ContractPriceAggregatorServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractPriceAggregatorServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceAggregatorServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractPriceAggregatorServiceManagerOwnershipTransferred)
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
func (it *ContractPriceAggregatorServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceAggregatorServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceAggregatorServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractPriceAggregatorServiceManager contract.
type ContractPriceAggregatorServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractPriceAggregatorServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractPriceAggregatorServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceAggregatorServiceManagerOwnershipTransferredIterator{contract: _ContractPriceAggregatorServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractPriceAggregatorServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractPriceAggregatorServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceAggregatorServiceManagerOwnershipTransferred)
				if err := _ContractPriceAggregatorServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractPriceAggregatorServiceManager *ContractPriceAggregatorServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractPriceAggregatorServiceManagerOwnershipTransferred, error) {
	event := new(ContractPriceAggregatorServiceManagerOwnershipTransferred)
	if err := _ContractPriceAggregatorServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
