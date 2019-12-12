package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Contact struct {
	Fullname string `json:"fullname"`
	Email map[string]string `json:"email"`
	Phone map[string]string `json:"phone"`
	OtherField map[string]string`json:"other_field"`
}

func generateNewContact(dataContact map[string]interface{}) Contact {
	contactPerson := make(map[string]string)
	otherField := make(map[string]string)
	c := dataContact["email"].(map[string]interface{})
	for k, v := range c {
		contactPerson[k] = v.(string)
	}

	contactPhone := make(map[string]string)
	cp := dataContact["phone"].(map[string]interface{})
	for k, v := range cp {
		contactPhone[k] = v.(string)
	}

	for k, v := range dataContact {
		if (k != "phone" && k != "fullname" && k != "email") {
			otherField[k] = v.(string)
		}
	}

	return Contact{
		Fullname: dataContact["fullname"].(string),
		Email: contactPerson,
		Phone: contactPhone,
		OtherField: otherField,
	}
}

func (c *Contact) appendData(dataContact map[string]interface{}) {
	contact := (*c)
	for key, value := range dataContact {
		switch key {
			case "fullname":
				continue
			case "email":
				var fields = value.(map[string]interface{})
				for k, v := range fields {
					contact.Email[k] = v.(string)
				}
			case "phone":
				var fields = value.(map[string]interface{})
				for k, v := range fields {
					contact.Phone[k] = v.(string)
				}
				break;
			default:
				contact.OtherField[key] = value.(string)
		}
	}
}

func MergeContact (resultJSON []map[string]interface{}) []Contact {
	var listContactByEmail = make(map[string]bool)
	var convertedContact []Contact
	var listMappedContact = make(map[string]Contact)

	for _, value := range resultJSON {
		var currentContactName = value["fullname"].(string)
		_, isEmailExist := listContactByEmail[currentContactName]

		if !isEmailExist {
			listContactByEmail[currentContactName] = true
			var newContact = generateNewContact(value)
			listMappedContact[currentContactName] = newContact
			convertedContact = append(convertedContact, newContact)
		} else {
			var selectedContact = listMappedContact[currentContactName]
			selectedContact.appendData(value)
		}
	}
	fmt.Println("CONTACT_MERGED")
	return convertedContact
}


func ParseJSONFile(fileName string) ([]map[string]interface{}, error){
	// Read JSON
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		panic(errors.New("JSON_PARSE_FAILED"))
		return nil, nil
	}

	fmt.Println("JSON_PARSED")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Merge The Contact
	var resultJSONParsed []map[string]interface{}
	json.Unmarshal([]byte(byteValue), &resultJSONParsed)
	return resultJSONParsed, nil
}

func main() {
	var resultJSON []map[string]interface{}
	resultJSON, _ = ParseJSONFile("mock/sample.json")

	var convertedContact = MergeContact(resultJSON)

	// Write Json To Response
	writeToJSON, _ := json.MarshalIndent(convertedContact, "", "	")
	_ = ioutil.WriteFile("response.json", writeToJSON, 0644)
}