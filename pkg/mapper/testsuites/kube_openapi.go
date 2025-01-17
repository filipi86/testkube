package testsuites

import (
	commonv1 "github.com/kubeshop/testkube-operator/apis/common/v1"
	testsuitesv1 "github.com/kubeshop/testkube-operator/apis/testsuite/v1"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

// MapTestSuiteListKubeToAPI maps TestSuiteList CRD to list of OpenAPI spec TestSuite
func MapTestSuiteListKubeToAPI(cr testsuitesv1.TestSuiteList) (tests []testkube.TestSuite) {
	tests = make([]testkube.TestSuite, len(cr.Items))
	for i, item := range cr.Items {
		tests[i] = MapCRToAPI(item)
	}

	return
}

// MapCRToAPI maps TestSuite CRD to OpenAPI spec TestSuite
func MapCRToAPI(cr testsuitesv1.TestSuite) (test testkube.TestSuite) {
	test.Name = cr.Name
	test.Namespace = cr.Namespace
	test.Description = cr.Spec.Description

	for _, s := range cr.Spec.Before {
		test.Before = append(test.Before, mapCRStepToAPI(s))
	}
	for _, s := range cr.Spec.Steps {
		test.Steps = append(test.Steps, mapCRStepToAPI(s))
	}
	for _, s := range cr.Spec.After {
		test.After = append(test.After, mapCRStepToAPI(s))
	}

	test.Description = cr.Spec.Description
	test.Repeats = int32(cr.Spec.Repeats)
	test.Labels = cr.Labels
	test.Schedule = cr.Spec.Schedule
	test.Variables = MergeVariablesAndParams(cr.Spec.Variables, cr.Spec.Params)
	test.Created = cr.CreationTimestamp.Time

	return
}

// mapCRStepToAPI maps CRD TestSuiteStepSpec to OpenAPI spec TestSuiteStep
func mapCRStepToAPI(crstep testsuitesv1.TestSuiteStepSpec) (teststep testkube.TestSuiteStep) {

	switch true {
	case crstep.Execute != nil:
		teststep = testkube.TestSuiteStep{
			StopTestOnFailure: crstep.Execute.StopOnFailure,
			Execute: &testkube.TestSuiteStepExecuteTest{
				Name:      crstep.Execute.Name,
				Namespace: crstep.Execute.Namespace,
			},
		}

	case crstep.Delay != nil:
		teststep = testkube.TestSuiteStep{
			Delay: &testkube.TestSuiteStepDelay{
				Duration: crstep.Delay.Duration,
			},
		}
	}

	return
}

// @Depracated
// MapDepratcatedParams maps old params to new variables data structure
func MapDepratcatedParams(in map[string]testkube.Variable) map[string]string {
	out := map[string]string{}
	for k, v := range in {
		out[k] = v.Value
	}
	return out
}

// MapCRDVariables maps variables between API and operator CRDs
// TODO if we could merge operator into testkube repository we would get rid of those mappings
func MapCRDVariables(in map[string]testkube.Variable) map[string]testsuitesv1.Variable {
	out := map[string]testsuitesv1.Variable{}
	for k, v := range in {
		out[k] = testsuitesv1.Variable{
			Name:  v.Name,
			Type_: string(*v.Type_),
			Value: v.Value,
		}
	}
	return out
}

func MergeVariablesAndParams(variables map[string]testsuitesv1.Variable, params map[string]string) map[string]testkube.Variable {
	out := map[string]testkube.Variable{}
	for k, v := range params {
		out[k] = testkube.NewBasicVariable(k, v)
	}

	for k, v := range variables {
		if v.Type_ == commonv1.VariableTypeSecret {
			out[k] = testkube.NewSecretVariable(v.Name, v.Value)
		}
		if v.Type_ == commonv1.VariableTypeBasic {
			out[k] = testkube.NewBasicVariable(v.Name, v.Value)
		}
	}

	return out
}
