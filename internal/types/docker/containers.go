package docker

type ContainerName string

const (
	Web          ContainerName = "street_ninja_web"
	Postgres     ContainerName = "db"
	Redis        ContainerName = "redis"
	CeleryWorker ContainerName = "celery-worker"
	CeleryBeat   ContainerName = "celery-beat"
	Bootstrap    ContainerName = "bootstrap"
)

// var Container = struct {
// 	Web          ContainerName
// 	Postgres     ContainerName
// 	Redis        ContainerName
// 	CeleryWorker ContainerName
// 	CeleryBeat   ContainerName
// 	Bootstrap    ContainerName
// }{
// 	Web:          ContainerName("street_ninja_web"),
// 	Postgres:     ContainerName("db"),
// 	Redis:        ContainerName("redis"),
// 	CeleryWorker: ContainerName("celery-worker"),
// 	CeleryBeat:   ContainerName("celery-beat"),
// 	Bootstrap:    ContainerName("bootstrap"),
// }
