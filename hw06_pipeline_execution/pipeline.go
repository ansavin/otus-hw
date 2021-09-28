package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	pipeIn := in

	for _, s := range stages {
		pipeOut := make(Bi)

		go func(s Stage, pipeOut Bi, pipeIn In) {
			defer close(pipeOut)
			ch := s(pipeIn)
			for {
				var data interface{}
				var ok bool
				select {
				case data, ok = <-ch:
					if !ok {
						return
					}
					pipeOut <- data
				case <-done:
					return
				}
			}
		}(s, pipeOut, pipeIn)
		pipeIn = pipeOut
	}
	return pipeIn
}
