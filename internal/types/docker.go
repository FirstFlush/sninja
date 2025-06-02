package types

type ContainerName string

var Container = struct {
	Web          ContainerName
	Postgres     ContainerName
	Redis        ContainerName
	CeleryWorker ContainerName
	CeleryBeat   ContainerName
	Bootstrap    ContainerName
}{
	Web:          ContainerName("web"),
	Postgres:     ContainerName("db"),
	Redis:        ContainerName("redis"),
	CeleryWorker: ContainerName("celery-worker"),
	CeleryBeat:   ContainerName("celery-beat"),
	Bootstrap:    ContainerName("bootstrap"),
}
