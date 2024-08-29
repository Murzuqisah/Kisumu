## Kisumu

To create the "kisumu" interpreted programming language in Go, here's an overview and breakdown of the requirements:

1. **Language Syntax**:
   - Similar to Go, but custom-designed for the kisumu language.
   - Files should have a `.ksm` extension.

2. **Data Structures**:
   - **Number**: Handle both integers and floats.
   - **String**: Manage string values.
   - **Boolean**: Represent true/false values.
   - **Null**: Represent the absence of value.
   - **Array**: List-like structure.
   - **Object/Hash**: Key-value pairs (dictionaries).

3. **Methods/Functions for Each Data Structure**:
   - **Number**: Methods like `add`, `subtract`, `multiply`.
   - **String**: Methods like `length`, `substring`, `concat`.
   - **Boolean**: Methods like `not`, `and`, `or`.
   - **Null**: Methods might be less applicable here but could include `isNull`, `isNotNull`.
   - **Array**: Methods like `length`, `first`, `last`, `get(index)`.
   - **Object/Hash**: Methods like `get(key)`, `set(key, value)`, `keys`.

4. **Implementation Steps**:
   - **Parser**: Read and parse `.ksm` files.
   - **Interpreter**: Execute the parsed code.
   - **Data Structures**: Define and implement each data structure and its methods.
   - **Error Handling**: Implement robust error handling for invalid operations or syntax.

5. **Extra Features** (optional):
   - Advanced features for the data structures or language constructs beyond the minimum requirements.


## Project Structure


## stack-based
- In a stack-based language, operations are performed using a stack data structure. Here's what it means:

1. **Stack Usage**: A stack is a Last In, First Out (LIFO) data structure. Operations are pushed onto the stack and popped off the stack. 

2. **Execution Model**: When executing code, values are pushed onto the stack and operations like addition or subtraction pop values off the stack, perform the operation, and push the result back onto the stack.

3. **Example**: For an expression like `3 4 +`, the values `3` and `4` are pushed onto the stack. The `+` operation pops these values, adds them, and pushes the result (`7`) back onto the stack.

4. **Language Design**: In a stack-based language, you would typically define functions and operations that manipulate the stack. For instance, a `print` function would pop the top value off the stack and output it.

5. **Advantages**: Stack-based languages can be simple to implement and efficient in terms of memory usage because they don't require explicit variable storage and management.

Integrating tools like a compiler, code formatter, dependency manager, test runner, scanner for coding mistakes, and linters into the development workflow for creating an interpreted language like kisumu involves several steps. Here’s how you can approach each aspect and integrate them into your workflow:

### 1. **Compiler**:

For an interpreted language, you don’t need a traditional compiler, but you do need a parser and an interpreter. The parser will translate kisumu code into an intermediate representation (like an Abstract Syntax Tree), and the interpreter will execute this representation. 

- **Implementation**: Implement a parser to process `.ksm` files and convert them into an AST. The interpreter will then walk the AST and execute the code.

### 2. **Code Formatter**:

A code formatter ensures that kisumu code is consistently formatted. While Go has `gofmt` for formatting Go code, you need to create or use a formatter for kisumu.

- **Implementation**: Write a code formatter that adheres to the style rules you define for kisumu. This formatter will analyze the syntax and output formatted code. You could also create a tool to integrate with your editor or IDE.

### 3. **Dependency Manager**:

In Go, `go mod` handles dependencies. For kisumu, you might not have dependencies in the same way but can manage libraries or modules if your interpreter or runtime has them.

- **Implementation**: If kisumu grows to include libraries, you could implement a simple dependency management system to download and manage these libraries, or use a Go module to manage the dependencies of the kisumu interpreter itself.

### 4. **Test Runner**:

To ensure your interpreter and language features work correctly, you need a test runner to execute tests and validate the functionality of your code.

- **Implementation**: Write test cases for your parser, interpreter, and other components. Use Go’s testing framework (`testing` package) to run these tests and report results.

### 5. **Scanner for Coding Mistakes**:

A static code analysis tool can help identify common coding mistakes or potential issues in kisumu code.

- **Implementation**: Develop a static analyzer that checks for common issues in kisumu code, such as syntax errors or logical mistakes. This can be integrated into your development workflow to provide feedback on code quality.

### 6. **Linters**:

Linters enforce coding standards and help catch errors before runtime.

- **Implementation**: Create or configure a linter that checks kisumu code against predefined coding standards. This could include style rules, best practices, and potential errors.

### **Integration into Workflow**:

1. **Development Environment**: Set up your development environment with tools for formatting, linting, and testing kisumu code. Integrate these tools into your editor or IDE for a seamless experience.

2. **Build Automation**: Use Makefiles or build scripts to automate the process of running tests, formatting code, and scanning for errors. This ensures consistency and simplifies running these tools.

3. **Continuous Integration**: Implement CI/CD pipelines using tools like GitHub Actions, GitLab CI, or Jenkins to automate testing, linting, and other checks for every code change. This helps maintain code quality and catches issues early.

4. **Documentation**: Document the usage of these tools and their integration into the development process in your project’s README or developer guide.

By integrating these tools effectively, you ensure a smooth development process and maintain high code quality for the kisumu interpreter and its ecosystem.