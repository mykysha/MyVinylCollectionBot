package pooler

type WorkerPool interface {
	Start()
	AddTask(task func())
	Stop()
}

type Pool struct {
	workerNumber int
	taskChan     chan func()
	stopChan     chan interface{}
}

func NewPool(workerNumber int) *Pool {
	taskChan := make(chan func(), workerNumber)
	stopChan := make(chan interface{})

	return &Pool{
		workerNumber: workerNumber,
		taskChan:     taskChan,
		stopChan:     stopChan,
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.workerNumber; i++ {
		go func() {
			for task := range p.taskChan {
				select {
				case <-p.stopChan:
					break
				default:
					task()
				}
			}
		}()
	}
}

func (p *Pool) AddTask(task func()) {
	p.taskChan <- task
}

func (p *Pool) Stop() {
	close(p.stopChan)
}
