package repository_test

import (
	"sync"
	"context"

	"algogrit.com/emp-server/employee/entities"
	"algogrit.com/emp-server/employee/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"syreclabs.com/go/faker"
)

var inmem = Describe("InmemRepository", func() {
	var inmemRepo = repository.NewInMem()

	Context("consistency", func() {
		Describe("Save", func() {
			var ctx, cancelFn  = context.WithCancel(context.Background())
			var wg sync.WaitGroup

			BeforeEach(func() {
				for i:= 0; i< 10; i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						emp := entities.Employee{
							Name: faker.Name().FirstName(),
							ProjectID: faker.Number().NumberInt(4),
							Designation: faker.RandomString(8),
						}
						<-ctx.Done()
						inmemRepo.Save(emp)
					}()
				}
			})

			It("should save a set of employees", func() {
				cancelFn()
				wg.Wait()
				Expect(len(inmemRepo.RetrieveAll())).To(Equal(12))
			})
		})
	})
})
