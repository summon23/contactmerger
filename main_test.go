package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseJSONFile(t *testing.T) {
	var fileName = "mock/sampleHappyFlow.json"
	resultParsed, _ := ParseJSONFile(fileName)

	assert.NotEmpty(t, resultParsed)
}

func TestMergeContact(t *testing.T) {
	var fileNameSample = "mock/sampleHappyFlow.json"
	sampleData, _ := ParseJSONFile(fileNameSample)
	resultSample := MergeContact(sampleData)

	var fileNameExpected = "mock/sampleResultHappyFlow.json"
	resultDataExpected, _ := ParseJSONFile(fileNameExpected)
	resultDataExpectedMerged := MergeContact(resultDataExpected)

	assert.EqualValues(t, resultSample, resultDataExpectedMerged, "DATA_NOT_MATCH_WITH_EXPECTED")
}