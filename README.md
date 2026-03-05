## myers-algo

`myers-algo` is a text differencing engine which calculates the differences between two files, using the **Myers Diff Algorithm**, it focuses on finding the **Shortest Edit Script (SES)** to transform one text into another, providing clear, color-coded terminal output.

### How it Works

The engine treats the diffing process as a graph search problem. By navigating a 2D grid where the axes represent the lines of the "old" and "new" files, the algorithm finds the optimal path from the top-left $(0,0)$ to the bottom-right $(N, M)$.

### Key Features

* **Myers Algorithm Implementation:** Utilizes the standard greedy approach for calculating differences, ensuring minimal edit sequences.
* **Backtracking Trace:** Implements a full trace history to reconstruct the edit sequence accurately after the optimal distance $D$ is found.
* **Terminal Color Support:** Built-in ANSI escape sequences for immediate visual feedback in the console.
* **Cross-Platform Normalization:** Automatically handles `\r\n` (Windows) and `\n` (Unix) line endings to ensure consistent diffs across environments.

### Getting Started

#### Prerequisites

* Go 1.21+

#### Installation

```bash
git clone https://github.com/ganimtron-10/myers-algo.git
cd myers-algo

```

#### Running the Tool

Ensure you have an `old.txt` and `new.txt` in the root directory, then execute:

```bash
go run main.go

```

### Example Output

If `old.txt` contains "Hello World" and `new.txt` contains "Hello Go", the output will look like:

```diff
- Hello World
+ Hello Go
```

### Project Structure

* `main.go`: Entry point, file I/O handling, and terminal output logic.
* `MyersAlgo`: The core logic that generates the $V$-array trace.
* `genEditSequence`: The backtracking function that converts the trace into a slice of `EditOperation` structs.
* `ComputeFileDiff`: The high-level wrapper that standardizes strings and produces the formatted string builder.
