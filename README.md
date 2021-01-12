<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Issues][issues-shield]][issues-url]
[![Apache 2.0 License][license-shield]][license-url]
[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.][status-shield]][status-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/sophiabrandt/go-bookclub">
    <img src="images/logo.png" alt="Logo">
  </a>

  <h3 align="center">Bookclub</h3>

  <p align="center">
    Go micro-service REST API
    <br />
    <a href="https://github.com/sophiabrandt/go-bookclub"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/sophiabrandt/go-bookclub">View Demo</a>
    ·
    <a href="https://github.com/sophiabrandt/go-bookclub/issues">Report Bug</a>
    ·
    <a href="https://github.com/sophiabrandt/go-bookclub/issues">Request Feature</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

![Books](images/alfons-morales-YLSwjSy7stw-unsplash.jpg)

_**Work in Progress**_

I'm learning Go. This is my attempt at creating a micro-service web application.

I'm following the [ardanlabs/service][service] video course to learn more about how to create idiomatic Go web applications, with inspirations from other resources (see [Acknowledgements](#acknowledgements)).

### Built With

- Go
- Docker

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these steps.

### Prerequisites

You'll need:

- Go > 1.11 (for Go module support)
- Docker, docker-compose
- [GNU make](https://www.gnu.org/software/make/)

### Installation

1. Clone the repo

   ```sh
   git clone https://github.com/sophiabrandt/go-bookclub.git
   ```

2. Use the Makefile for running commands:
   ```sh
   # tidies and vendors dependencies
   make tidy
   ```

I'm using Linux to develop this application - no guarantees for cross-platform compatibility. If you can't get `make` working, try to copy the necessary commands from the [`Makefile`](Makefile).

<!-- USAGE EXAMPLES -->

## Usage

coming soon

<!-- ROADMAP -->

## Roadmap

See the [open issues](https://github.com/sophiabrandt/go-bookclub/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

Please update tests if necessary.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under the Apache 2.0 License. See [`LICENSE`](LICENSE) for more information.

<!-- CONTACT -->

## Contact

Sophia Brandt - [@hisophiabrandt](https://twitter.com/hisophiabrandt)

Project Link: [https://github.com/sophiabrandt/go-bookclub](https://github.com/sophiabrandt/go-bookclub)

<!-- ACKNOWLEDGEMENTS -->

## Acknowledgements

- [ardanlabs][service]
- [ashishjuyal](https://github.com/ashishjuyal/banking)
- [dumindu](https://github.com/learning-cloud-native-go/myapp)
- [dlsniper](https://github.com/dlsniper/gopherconuk)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/sophiabrandt/go-bookclub.svg
[contributors-url]: https://github.com/sophiabrandt/go-bookclub/graphs/contributors
[issues-shield]: https://img.shields.io/github/issues/sophiabrandt/go-bookclub.svg
[issues-url]: https://github.com/sophiabrandt/go-bookclub/issues
[license-shield]: https://img.shields.io/github/license/sophiabrandt/go-bookclub.svg
[license-url]: https://github.com/sophiabrandt/go-bookclub/blob/master/LICENSE
[status-shield]: https://www.repostatus.org/badges/latest/wip.svg
[status-url]: https://www.repostatus.org/#wip
[service]: https://github.com/ardanlabs/service
