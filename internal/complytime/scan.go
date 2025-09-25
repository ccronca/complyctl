// SPDX-License-Identifier: Apache-2.0

package complytime

import (
	"encoding/json"
	"os"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

// WriteAssessmentResults writes AssessmentResults as a JSON file to a given path location.
func WriteAssessmentResults(assessmentResults *oscalTypes.AssessmentResults, assessmentResultsLocation string) error {

	oscalModels := oscalTypes.OscalModels{
		AssessmentResults: assessmentResults,
	}

	assessmentResults.Metadata.Title = replaceString(assessmentResults.Metadata.Title, "Assessment result")

	assessmentResultsJson, err := json.MarshalIndent(oscalModels, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(assessmentResultsLocation, assessmentResultsJson, 0600)

}
