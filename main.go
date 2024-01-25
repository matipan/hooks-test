package main

type Hooks struct{}

// example usage: "dagger call container-echo --string-arg yo"
func (m *Hooks) ContainerEcho(stringArg string) *Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}
