# 📗 Table of Contents

- [📖 About the Project](#about-project)
  - [🛠 Built With](#built-with)
    - [Tech Stack](#tech-stack)
    - [Key Features](#key-features)
- [💻 Getting Started](#getting-started)
  - [Setup](#setup)
  - [Prerequisites](#prerequisites)
  - [Install](#install)
  - [Usage](#usage)
- [👥 Authors](#authors)
- [🔭 Future Features](#future-features)
- [⭐️ Show your support](#support)
- [🙏 Acknowledgements](#acknowledgements)
- [📝 License](#license)

<!-- PROJECT DESCRIPTION -->

# 📖 [CC_Collaboration_server] <a name="about-project"></a>

> This Go application sets up a WebSocket server to enable real-time collaboration on a code editor. It utilizes the Gorilla WebSocket library to manage WebSocket connections and facilitate communication between clients.

## 🛠 Built With <a name="built-with"></a>

### Tech Stack <a name="tech-stack"></a>

<details>
  <summary>Server</summary>
  <ul>
    <li><a href="#">Golang</a></li>
    <li><a href="#">Gin Framework</a></li>
    <li><a href="#">Websockets</a></li>
  </ul>
</details>

<!-- Features -->

### Key Features <a name="key-features"></a>

- **WebSocket Communication**: Utilizes WebSocket technology for real-time, bidirectional communication between the server and connected clients.
- **Cross-Origin Resource Sharing (CORS)**: Implements CORS to allow cross-origin requests, enabling collaboration across different domains.
- **Simple Message Structure**: Defines a simple message structure for exchanging code content and file extensions between clients and the server.
- **Efficient Broadcasting**: Broadcasts messages efficiently to all connected clients, ensuring synchronized collaboration.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## 💻 Getting Started <a name="getting-started"></a>

To get a local copy up and running, follow these steps.

### Prerequisites

In order to run this project you need:

<ul>
    <li>Go</li>
    <li>Gin Framework</li>
    <li>Internet</li>
</ul>

### Setup

Clone this repository to your desired folder:

Example commands:

```bash
cd my-folder
git clone https://github.com/Rishi-Mishra0704/cc_editor_server
```

### Usage

To run the project, execute the following command:

```bash
cd cc_editor_server
go get ./...
```

```bash
go run main.go
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- AUTHORS -->

## 👥 Authors <a name="authors"></a>

👤 **Rishi_Mishra**

- GitHub: [@githubhandle](https://github.com/Rishi-Mishra0704)
- Twitter: [@twitterhandle](https://twitter.com/RishiMi31357764)
- LinkedIn: [LinkedIn](https://www.linkedin.com/in/rrmishra/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- SUPPORT -->

## ⭐️ Show your support <a name="support"></a>

> Show your support by giving a ⭐️ if you like this project!

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGEMENTS -->

## 🙏 Acknowledgments <a name="acknowledgements"></a>

I would like to thank HackSwift for giving me the opportunity to work on this project.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## 📝 License <a name="license"></a>

This project is [MIT](./LICENSE) licensed.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
