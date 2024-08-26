# Plush Database 🍥

## A Database Implementation Using B+ Trees in Golang

This project is an exploration into the internals of database systems, with a focus on understanding and implementing the core components of a database using B+ trees as the primary data structure. The goal is to create a simple, in-memory key-value store that supports basic CRUD operations.

### Project Structure

- **cmd/**: 
  - Contains the main application entry point (`main.go`), where the database is initialized and the command-line interface (CLI) is managed.

- **internal/**:
  - Holds internal packages that encapsulate the core functionality, including:
    - **bplustree/**: Implements the B+ tree logic, including node structures, insertion, deletion, and searching algorithms.
    - **storage/**: Manages data storage, including in-memory and file-based storage options.

- **pkg/**:
  - May contain reusable utility packages or libraries that could be shared across different projects or services.

- **data/**:
  - Used for data files, including any sample or test datasets, and potential data export/import functionality for the database.

### Timeline & Documentation

1. **Understanding B+ Trees**
   - Start with an in-depth study of B+ trees, focusing on their application as a balanced data structure for efficient querying, insertion, and deletion operations.
   - Implement the basic B+ tree structure, including:
     - Node structure (internal and leaf nodes)
     - Insertion algorithm
     - Deletion algorithm
     - Search operation

2. **Schema Design**
   - Design a flexible schema system that allows users to define the structure of their data.
   - The schema should support various data types, including:
     - **Primitive types**: `int`, `string`, `bool`
     - **Composite types**: `struct`, where a block of data or a node can hold multiple fields
   - Allow users to specify the number of fields (columns) and their types, which will guide how data is inserted and stored.

   **Example Schema Input:**
   - `{key: int, name: string, email: string}`

   **Example Node Structure:**
   ```go
   type Node struct {
       PrimaryKey int             // Used as the key in the B+ tree
       Data       map[string]interface{} // Holds the actual data, e.g., name, email
       Left       *Node           // Pointer to the left child (if applicable)
       Right      *Node           // Pointer to the right child (if applicable)
   }
   ```

3. **Implementing an In-Memory Database**
   - Develop the core database functionalities:
     - **Define Schema**: Allow users to define and update the schema dynamically.
     - **Add Data**: Insert new records into the database, ensuring they are stored according to the defined schema and indexed using the B+ tree.
     - **Remove Data**: Implement deletion of records, maintaining the integrity and balance of the B+ tree.
     - **Update Data**: Allow modifications to existing records.
     - **Query Data**: Retrieve records based on primary keys or other criteria.
     - **Get Size of B+ Tree**: Provide a function to return the number of nodes in the B+ tree.
     - **Get Total Size of Table**: Calculate the total amount of data stored in the table, including metadata.

4. **Interface Layers in Databases**
   - **SQL Layer**: Simulate a basic SQL-like interface for interacting with the database, supporting commands such as `SELECT`, `INSERT`, `UPDATE`, and `DELETE`.
   - **KV (Key-Value) Layer**: Implement a direct key-value interface for lower-level data access, bypassing the SQL layer for performance-critical operations.

5. **Advanced Features (Future Work)**
   - **Persistence**: Extend the in-memory database to support persistent storage, allowing data to be saved to disk and reloaded on startup.
   - **Concurrency**: Implement basic concurrency control to handle multiple simultaneous operations, ensuring data consistency.
   - **Indexing**: Explore the addition of secondary indexes for faster querying on non-primary key fields.

6. **Testing and Benchmarking**
   - Implement a comprehensive suite of unit tests to ensure the correctness of the B+ tree operations and database functionalities.
   - Benchmark the performance of various operations (e.g., insertion, deletion, querying) to identify bottlenecks and optimize the implementation.

7. **Documentation**
   - Write detailed documentation covering the project structure, usage, API endpoints (if applicable), and examples of how to interact with the database.
