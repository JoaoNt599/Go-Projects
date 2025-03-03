# Technical Assitance

In the domain of an application for a Technical Assistance Store, the occurrence
of the following entities: Technician, Attendant, Customer, Service, Service Request and Material.

The Customer calls the technical assistance store and requests a service.

The Attendant is the person who assists the customer and registers the service request. The Attendant defines who will be the
technician who will serve the customer (Service assignment) and records the assignment in the service request.

The Technician records the start date of the service, performs the service and records the completion date.
When the technician uses some material (products) to perform the service, he also records it in the request
of service which materials were used. This way, the customer can know who performed the service, when the service was
carried out and what materials were used.

When using the application, both the attendant and the technician use the system to record information.
The technician has the following characteristics: name, social security number, telephone number. The Customer has the characteristics name, CPF, address and
telephone.

The Attendant has the characteristics name and CPF. The Service Request has the following characteristics: number of
identification, request date, start date, end date, in addition it must present who the client is, who the
the technician and what materials were used.

The Service Request entity can inform the date of the request, who the customer is, who the attendant was, who was
the technician who performed the service and also what materials were used. Service has the following characteristics:
identification, description and value. Material has the characteristics of identification number, description and value.

All entities must have a behavior called “exibeInformacoes” to show internal data.
