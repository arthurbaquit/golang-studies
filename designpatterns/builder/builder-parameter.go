package main

import "fmt"

// The builder pattern should hide the real object from the cliente
// and should provide a way to build the object step by step. Therefore,
// the client should only interact with the builder object and not with
// the real object. An example of this is the EmailBuilder object:

type email struct {
	From, To, Subject, Body string
}

// notice that the email struct is not exported, as the client will never
// interact with it directly.

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	b.email.From = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.To = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.Subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.Body = body
	return b
}

func SendEmailImpl(email *email) {
	// send email
	fmt.Println("Sending email...")
	fmt.Println("From:", email.From)
	fmt.Println("To:", email.To)
	fmt.Println("Subject:", email.Subject)
	fmt.Println("Body:", email.Body)

}

type build func(*EmailBuilder)

// The SendEmail function is the only function that the client will
// interact with. It receives a function as a parameter, which will
// be used to build the email object. Therefore, the client will
// never interact with the email object directly or with its implementation.
func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	SendEmailImpl(&builder.email)
}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar.com").
			To("baz@bar.com").
			Subject("Hello world").
			Body("This is my builder pattern example")
	})
}
