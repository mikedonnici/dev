# UML

Unified Modelling Laguage provides a standardised diagramming technique for expressing the design of software developed using object-oriented principles.

## Software development methodologies

- Necessary for organisation as complexity and team size increases

### Waterfall Method

- Used when requirement sare clear and won't change frequently
- Linear model
- Requires pre-defined requirements and detailed plan
- No changes expected during development
- Used for life-control, medical and military systems

#### Phases

1. Collect an analyse requirements
   - Clarify with stakeholders
   - Document thoroughly
1. Architecture definition
   - Packages
   - Key types
   - Interactions
   - Security
   - Fault tolerance... etc
1. Implementation
   - Coding
   - Testing
1. Verification (testing)
   - Functionality
   - Performance
   - Usability
   - Security
   - Bug reporting and fix cycle
1. Maintenance
   - Fixing smaller bugs
   - Small enhancements

- Major changes require a new _waterfall_

### Agile Software Development

- Change-friendly and flexible
- A way of thinking rather than a strict methodology - an approach
- Began with [Agile Manifesto](https://agilemanifesto.org/) in 2001 which aimed to end the proliferation of software development methodologies
- Agile manifesto defines 4 values:
  1. Individuals and interactions over processes and tools
  2. Working software over comprehensive documentation
  3. Customer collaboration over contract negotiation
  4. Responding to change over a plan
- Still requires tools and processes - but adaptive
- Use documentation when it provides real value
- Still need contracts to manage expectations and agreed outcomes, but with a focus on partnership, collaboration and trust
- Planning required but not at a level of detail that creates blocks to progress
- Iterative delivery
- Work broken in smaller chunks (sprints) generally 2-4 weeks long
- Review at the end of each sprint
- Testing toghtly integrated with development
- Business users and stakeholders closely involved

## Object-oriented concepts

### Objects

- Describe things (nouns) in terms of:
  - properties
  - identity (own state)
  - behaviour

### Classes

- Provides a blueprint for an object:
  - name
  - properties
  - methods

### Abstraction

- Describing complex problems, in simple terms, by ignoring some details
- When defining a class, focuss on essential details and ignore the rest

### Encapsulation and Data Hiding

- Packing attributes and behaviour into a class
- Hiding internals and only exposing what is necessary for use
- Protects object from undesirable changes
- Protects classes from unwanted external dependancies (tightly coupled - bad design)

### Inheritence

- Allows the re-use of class code in new classes
- Subclasses inherit from superclasses
- Facilitates DRY principle and polymorphism

### Polymorphism

- Means occuring in more than one form
- Inherited attributes can be overridden (method overriding) so that subclass object can behave differently from its superclass

## Object-oriented analysis and design

1. [Collect requirements](#collect-requirements):
   1. **Identify** the _problems_ we want to solve
   1. **Clarify** the _functionality_ required to solve the problems
   1. **Document** important _decisions_ that are reached, as clearly as possible
1. [Describe the software](#describe-the-software):
   1. **Describe** the system from the _user's_ perspective
   1. For agile, don't need a lot of detail
   1. Create wireframes, non-functional prototypes if needed
1. [Identify the classes](#identify-the-classes):
   1. Should be easy if previous steps were well executed
   1. Effectively a set of significant nouns desribing things and services
1. Create diagrams:
   1. Visual representations of classes, their attributes and behaviour
   1. Model interactions between objects

### Collect requirements

Define what the system must do. Identify the constraints and boundaries.

- Often called _requirements analysis_
- Paves the way for the rest of the OO design process
- A _requirement_ is something that is needed or wanted
- **Functional requirements**:
  - software features
  - reactions to input
  - Expected behaviour
- **Non-functional requirements**:

  - performance / availability
  - security
  - legal considerations
  - documentation

Documenting requirements can be done is simple, short phrases in the form:

"The app/system must \<do something\>"

For _waterfall_ this step must be complete, for agile requirements can be refined along the way.

### Describe the software

**Mapping requirements to technical descriptions**

This is where we describe the systems _from the user's point of view_.

A couple of methods:

- **Use cases**:

  - Each describes a distinct piece of functionality
  - _Title_: short, descriptive
  - _Actor_: user or other system
  - _Scenario_: Describes how the system works in this case, bullet points - 'The user can..." is a good format
  - Avoid technical terms, should be understood by all stakeholders

- **User stories**:

  - Very brief description of a feature, usually only one or two sentences
  - Format is often: _"As a \<type of user\> I want to \<some goal\> so that \<some reason\>"_

- **Epics**:
  - A larger or more top-level _user story_ cannot be expressed simply in one or two sentences and is called an _epic_
  - Epics must be further divided into smaller user stories

_User stories_ might be captured on sticky notes or cards and are not intended to capture the feature details. Rather, they are used to prime discussions. They are about communication and are commonly used in agile development.

_Use cases_ are used in waterfall methodologies.

### Identify the classes

...