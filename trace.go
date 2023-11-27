package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type Trace struct {
}

type Compound_load struct {
	Trace_code         string
	CompoundID         string
	Timestamp          string
	Manufacturer       string
	Responsible_person string
	Report_hash        string
}

type Storage_load struct {
	Trace_code         string
	Location           string
	Source             string
	CompoundID         string
	Timestamp          string
	Responsible_person string
	Report_hash        string
}

type Transport_load struct {
	Trace_code         string
	Source             string
	CompoundID         string
	Timestamp          string
	Destination        string
	Responsible_person string
	Report_hash        string
}

type Expend_load struct {
	Trace_code         string
	Source             string
	CompoundID         string
	Timestamp          string
	Responsible_person string
	Operator           string
	Report_hash        string
}

func (t *Trace) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *Trace) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "record_compund":
		return t.record_compund(stub, args)
	case "record_storage":
		return t.record_storage(stub, args)
	case "record_transport":
		return t.record_transport(stub, args)
	case "record_expend":
		return t.record_expend(stub, args)
	case "query":
		result, err := t.query(stub, args)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(result))
	}
	return shim.Error("Invalid invoke function of " + function)
}

func (t *Trace) record_compund(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 6 {
		return shim.Error("Invalid args length : " + strconv.Itoa(len(args)) + "required args length : 6")
	}
	//将参数填入数据结构中
	uploadData := &Compound_load{
		Trace_code:         args[0],
		CompoundID:         args[1],
		Timestamp:          args[2],
		Manufacturer:       args[3],
		Responsible_person: args[4],
		Report_hash:        args[5],
	}

	//将对象序列化成数组存储
	newdata, _ := json.Marshal(uploadData)
	//所有数据根据Trace_code查询，Trace_code作为键，值为多个交易数据的集合
	//先取出键对应的原值，再添加新交易数据
	existingValue, err := stub.GetState(uploadData.Trace_code)
	if err != nil {
		return shim.Error("Failed to get state: " + err.Error())
	}
	newValue := append(existingValue, newdata...)
	err = stub.PutState(uploadData.Trace_code, newValue)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *Trace) record_storage(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 7 {
		return shim.Error("Invalid args length : " + strconv.Itoa(len(args)) + "required args length : 7")
	}

	uploadData := &Storage_load{
		Trace_code:         args[0],
		Location:           args[1],
		Source:             args[2],
		CompoundID:         args[3],
		Timestamp:          args[4],
		Responsible_person: args[5],
		Report_hash:        args[6],
	}

	newdata, _ := json.Marshal(uploadData)
	existingValue, err := stub.GetState(uploadData.Trace_code)
	if err != nil {
		return shim.Error("Failed to get state: " + err.Error())
	}
	newValue := append(existingValue, newdata...)
	err = stub.PutState(uploadData.Trace_code, newValue)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *Trace) record_transport(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 7 {
		return shim.Error("Invalid args length : " + strconv.Itoa(len(args)) + "required args length : 7")
	}

	uploadData := &Transport_load{
		Trace_code:         args[0],
		Source:             args[1],
		CompoundID:         args[2],
		Timestamp:          args[3],
		Destination:        args[4],
		Responsible_person: args[5],
		Report_hash:        args[6],
	}

	newdata, _ := json.Marshal(uploadData)
	existingValue, err := stub.GetState(uploadData.Trace_code)
	if err != nil {
		return shim.Error("Failed to get state: " + err.Error())
	}
	newValue := append(existingValue, newdata...)
	err = stub.PutState(uploadData.Trace_code, newValue)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *Trace) record_expend(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 7 {
		return shim.Error("Invalid args length : " + strconv.Itoa(len(args)) + "required args length : 7")
	}

	uploadData := &Expend_load{
		Trace_code:         args[0],
		Source:             args[1],
		CompoundID:         args[2],
		Timestamp:          args[3],
		Responsible_person: args[4],
		Operator:           args[5],
		Report_hash:        args[6],
	}

	newdata, _ := json.Marshal(uploadData)
	existingValue, err := stub.GetState(uploadData.Trace_code)
	if err != nil {
		return shim.Error("Failed to get state: " + err.Error())
	}
	newValue := append(existingValue, newdata...)
	err = stub.PutState(uploadData.Trace_code, newValue)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *Trace) query(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(value), nil
}

func main() {
	if err := shim.Start(new(Trace)); err != nil {
		fmt.Printf("Error starting Trace chaincode:%s", err)
	}
}
