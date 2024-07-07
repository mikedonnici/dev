# General Dev Notes

## Types of software tests

1. **Unit Tests**: 
   - Testing individual units or components of the software in isolation.
   - Isolated from external dependencies using mocks or stubs.
   - Tests individual functions or methods.
2. **Component Tests**:
   - Testing a group of related units together to achieve a specific functionality.
   - Generally isolated from external dependencies.
   - Tests a set of functions or methods such as a class or module.
   - May involve use of a test (docker) database or similar.
   - If required, should mock external systems (eg external APIs)
3. **Integration Testing**:
    - Testing the integration of multiple components or services.
    - Tests the interaction between multiple components or services.
    - May involve use of a test (docker) database or similar. 
    - May mock external systems or test the _real_ interaction.
4. **End-to-End Testing**:
    - Testing the entire system from end to end.
    - Uses _real_ environment and _real_ external dependencies. 
