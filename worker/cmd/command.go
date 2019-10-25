package cmd

import (
	"github.com/khorevaa/go-AutoUpdate1C/cmd/flag"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/sigmon"
	"os"
	"time"
)

type WorkerConfig struct {
	Name     string   `long:"name"  description:"The name to set for the worker during registration. If not specified, the hostname will be used."`
	Tags     []string `long:"tag"   description:"A tag to set during registration. Can be specified multiple times."`
	TeamName string   `long:"team"  description:"The name of the team that this worker will be assigned to."`

	Version string `long:"version" hidden:"true" description:"Version of the worker. This is normally baked in to the binary, so this flag is hidden."`
}

type TSAConfig struct {
	Hosts     []string `long:"host" default:"127.0.0.1:2222" description:"TSA host to forward the worker through. Can be specified multiple times."`
	WorkerSid string   `long:"sid" description:"Worker sid expect from the TSA."`
}

type WorkerCommand struct {
	Worker WorkerConfig

	TSA                TSAConfig     `group:"TSA Configuration" namespace:"tsa"`
	WorkDir            flag.Dir      `long:"work-dir" required:"true" description:"Directory in which to place work data."`
	HealthCheckTimeout time.Duration `long:"healthcheck-timeout"    default:"5s"       description:"HTTP timeout for the full duration of health checking."`
}

func (cmd *WorkerCommand) Execute(args []string) error {
	runner, err := cmd.Runner(args)
	if err != nil {
		return err
	}

	return <-ifrit.Invoke(sigmon.New(runner)).Wait()
}

func (cmd *WorkerCommand) Runner(args []string) (ifrit.Runner, error) {

	//atcWorker, gardenRunner, err := cmd.gardenRunner(logger.Session("garden"))
	//if err != nil {
	//	return nil, err
	//}
	//
	//atcWorker.Version = concourse.WorkerVersion
	//
	//baggageclaimRunner, err := cmd.baggageclaimRunner(logger.Session("baggageclaim"))
	//if err != nil {
	//	return nil, err
	//}
	//
	//healthChecker := worker.NewHealthChecker(
	//	logger.Session("healthchecker"),
	//	cmd.baggageclaimURL(),
	//	cmd.gardenURL(),
	//	cmd.HealthCheckTimeout,
	//)
	//
	//tsaClient := cmd.TSA.Client(atcWorker)
	//
	//beaconRunner := worker.NewBeaconRunner(
	//	logger.Session("beacon-runner"),
	//	tsaClient,
	//	cmd.RebalanceInterval,
	//	cmd.ConnectionDrainTimeout,
	//	cmd.gardenAddr(),
	//	cmd.baggageclaimAddr(),
	//)
	//
	//gardenClient := gclient.BasicGardenClientWithRequestTimeout(
	//	logger.Session("garden-connection"),
	//	atc.GARDEN_CLIENT_HTTP_TIMEOUT,
	//	cmd.gardenURL(),
	//)
	//
	//baggageclaimClient := bclient.NewWithHTTPClient(
	//	cmd.baggageclaimURL(),
	//
	//	// ensure we don't use baggageclaim's default retryhttp client; all
	//	// traffic should be local, so any failures are unlikely to be transient.
	//	// we don't want a retry loop to block up sweeping and prevent the worker
	//	// from existing.
	//	&http.Client{
	//		Transport: &http.Transport{
	//			// don't let a slow (possibly stuck) baggageclaim server slow down
	//			// sweeping too much
	//			ResponseHeaderTimeout: 1 * time.Minute,
	//		},
	//		// we've seen destroy calls to baggageclaim hang and lock gc
	//		// gc is periodic so we don't need to retry here, we can rely
	//		// on the next sweeper tick.
	//		Timeout: 5 * time.Minute,
	//	},
	//)
	//
	//containerSweeper := worker.NewContainerSweeper(
	//	logger.Session("container-sweeper"),
	//	cmd.SweepInterval,
	//	tsaClient,
	//	gardenClient,
	//	cmd.ContainerSweeperMaxInFlight,
	//)
	//
	//volumeSweeper := worker.NewVolumeSweeper(
	//	logger.Session("volume-sweeper"),
	//	cmd.SweepInterval,
	//	tsaClient,
	//	baggageclaimClient,
	//	cmd.VolumeSweeperMaxInFlight,
	//)

	var members grouper.Members

	//if !cmd.gardenIsExternal() {
	//	members = append(members, grouper.Member{
	//		Name:   "garden",
	//		Runner: concourseCmd.NewLoggingRunner(logger.Session("garden-runner"), gardenRunner),
	//	})
	//}

	members = append(members, grouper.Members{
		//{
		//	Name:   "baggageclaim",
		//	Runner: concourseCmd.NewLoggingRunner(logger.Session("baggageclaim-runner"), baggageclaimRunner),
		//},
		//{
		//	Name: "debug",
		//	Runner: concourseCmd.NewLoggingRunner(
		//		logger.Session("debug-runner"),
		//		http_server.New(
		//			fmt.Sprintf("%s:%d", cmd.DebugBindIP.IP, cmd.DebugBindPort),
		//			http.DefaultServeMux,
		//		),
		//	),
		//},
		//{
		//	Name: "healthcheck",
		//	Runner: concourseCmd.NewLoggingRunner(
		//		logger.Session("healthcheck-runner"),
		//		http_server.New(
		//			fmt.Sprintf("%s:%d", cmd.HealthcheckBindIP.IP, cmd.HealthcheckBindPort),
		//			http.HandlerFunc(healthChecker.CheckHealth),
		//		),
		//	),
		//},
		//{
		//	Name: "beacon",
		//	Runner: concourseCmd.NewLoggingRunner(
		//		logger.Session("beacon-runner"),
		//		beaconRunner,
		//	),
		//},
		//{
		//	Name: "container-sweeper",
		//	Runner: concourseCmd.NewLoggingRunner(
		//		logger.Session("container-sweeper"),
		//		containerSweeper,
		//	),
		//},
		//{
		//	Name: "volume-sweeper",
		//	Runner: concourseCmd.NewLoggingRunner(
		//		logger.Session("volume-sweeper"),
		//		volumeSweeper,
		//	),
		//},
	}...)

	return grouper.NewParallel(os.Interrupt, members), nil
}
