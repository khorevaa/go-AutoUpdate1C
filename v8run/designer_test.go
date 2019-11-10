package v8run

import (
	"github.com/khorevaa/go-AutoUpdate1C/v8run/types"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"testing"
)

type designerTestSuite struct {
	baseTestSuite
	tempIB types.InfoBase
	v8path string
	ibPath string
}

func TestDesigner(t *testing.T) {
	suite.Run(t, new(designerTestSuite))
}

func (t *designerTestSuite) SetupSuite() {
	t.v8path = "/opt/1cv8/8.3.15.1194/1cv8"
	ibPath, _ := ioutil.TempDir("", "1c_DB_")
	t.ibPath = ibPath
}

func (t *designerTestSuite) AfterTest(suite, testName string) {
	t.clearTempInfoBase()
}

func (t *designerTestSuite) BeforeTest(suite, testName string) {
	t.createTempInfoBase()
}

func (t *designerTestSuite) TearDownTest() {

}

func (t *designerTestSuite) createTempInfoBase() {

	ib := NewFileIB(t.ibPath)

	_, err := Run(ib, CreateTempInfoBase(),
		WithPath(t.v8path),
		WithTimeout(30))

	t.tempIB = ib

	t.r().NoError(err)

}

func (t *designerTestSuite) clearTempInfoBase() {

	err := os.RemoveAll(t.ibPath)
	t.r().NoError(err)
}

func (t *designerTestSuite) TestLoadCfg() {

	_, err := Run(t.tempIB, LoadCfg("/Users/khorevaa/GolandProjects/go-AutoUpdate1C/tests/fixtures/0.9/1Cv8.cf"),
		WithPath(t.v8path),
		WithTimeout(30),
		WithOut("/Users/khorevaa/GolandProjects/go-AutoUpdate1C/v8run/log.txt", false))

	t.r().NoError(err)

	//t.r().Equal(len(codes), 1, "Промокод должен быть START")
	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")

}

func (t *designerTestSuite) TestUserOptions() {

	fnLoadCfg := func(file string) types.UserOption {
		return func(o types.Optioned) {
			o.SetOption("/LoadCfg", file)
		}
	}

	fnUpdateDBCfg := func() types.UserOption {
		return func(o types.Optioned) {
			o.SetOption("/UpdateDBCfg", true)
		}
	}

	task := NewDesigner(fnLoadCfg("/Users/khorevaa/GolandProjects/go-AutoUpdate1C/tests/fixtures/0.9/1Cv8.cf"),
		fnUpdateDBCfg())

	_, err := Run(t.tempIB, task,
		WithPath(t.v8path),
		WithTimeout(30),
		WithOut("/Users/khorevaa/GolandProjects/go-AutoUpdate1C/v8run/log.txt", false))

	t.r().NoError(err)

	//t.r().Equal(len(codes), 1, "Промокод должен быть START")
	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")

}

func (t *designerTestSuite) TestLoadCfgWithUpdateCfgDB() {

	loadCfg := LoadCfg("/Users/khorevaa/GolandProjects/go-AutoUpdate1C/tests/fixtures/0.9/1Cv8.cf")
	loadCfg.WithUpdateDBCfg(UpdateDBCfg(false, false))

	_, err := Run(t.tempIB, loadCfg,
		WithPath(t.v8path),
		WithTimeout(30),
		WithOut("/Users/khorevaa/GolandProjects/go-AutoUpdate1C/v8run/log.txt", false))

	t.r().NoError(err)

	//t.r().Equal(len(codes), 1, "Промокод должен быть START")
	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")

}

func (t *designerTestSuite) TestUpdateCfg() {

	loadCfg := LoadCfg("/Users/khorevaa/GolandProjects/go-AutoUpdate1C/tests/fixtures/0.9/1Cv8.cf")
	loadCfg.WithUpdateDBCfg(UpdateDBCfg(false, false))
	_, err := Run(t.tempIB, loadCfg,
		WithPath(t.v8path),
		WithTimeout(30),
		WithOut("/Users/khorevaa/GolandProjects/go-AutoUpdate1C/v8run/log.txt", true))

	t.r().NoError(err)

	task := UpdateCfg("/Users/khorevaa/GolandProjects/go-AutoUpdate1C/tests/fixtures/1.0/1Cv8.cf", false)
	task.WithUpdateDBCfg(UpdateDBCfg(false, false))

	_, err = Run(t.tempIB, task,
		WithPath(t.v8path),
		WithTimeout(30),
		WithOut("/Users/khorevaa/GolandProjects/go-AutoUpdate1C/v8run/log.txt", true))

	t.r().NoError(err)

	//t.r().Equal(len(codes), 1, "Промокод должен быть START")
	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")

}
