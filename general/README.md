# General

## Continuous Integration, Delivery and Deployment

**Continuous Integration** infers the integration of code updates into the main repository
multiple times per day. Development tasks need to be small enough for
code to be committed regularly.

Larger tasks that might require multiple days between commits might be
referred to as *frequent integration*.

My def: **CODE IS COMMITTED TO MAIN LINE*

**Continuous Delivery** has two main schools of thought. The first,
that it is an extension of the term *continuous integration*.

The second, from the opening line in the [Agile Manfesto](http://agilemanifesto.org/principles.html):

> Our highest priority is to satisfy the customer through early and
> continuous delivery of valuable software.

Thus, **continuos delivery** would include everything required to
automatically verify that a task is [done](https://www.agilealliance.org/glossary/definition-of-done).
This is analagos to DevOps.

My def: **CODE TASK IS DONE, COMMITTED AND VERIFIED BY AUTOMATED TESTS**

**Continuous Deployment** is often used as a synomym for *continuous delivery*
or may infer the additional step of the actual deployment of the updated
code into the production environment.

My def: **AS ABOVE WITH DEPLOYMENT TO END USERS**

## Refs

* https://stackoverflow.com/questions/28608015/continuous-integration-vs-continuous-delivery-vs-continuous-deployment
* http://agilemanifesto.org/principles.html
* https://www.agilealliance.org/
