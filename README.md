<a name="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/hood/coso">
    <img src="https://i.imgur.com/hBruQVEm.jpg" alt="Logo" width="180" height="180">
  </a>

<h3 align="center">coso</h3>

  <p align="center">
    An open-source HTTP back-end with realtime subscriptions using Google Cloud Storage as a key-value store.
    <br />
    <a href="https://github.com/hood/coso/README.md"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/hood/coso/issues">Report Bug</a>
    ·
    <a href="https://github.com/hood/coso/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#building-and-running-coso">
        Building and running coso
      </a>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

`coso` is an **experimental** open-source HTTP back-end using Google Cloud Storage
as a key-value store. It is designed to be a simple and fast way to use Google
Cloud Storage as a key-value store with a HTTP interface. Real-time
subscriptions are also possible via Server Sent Events.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Building and running coso

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

```sh
git clone https://github.com/hood/coso
cd coso
go build
./coso
```

`coso` will look for an `.env` file in its own folder, and will use the following environment variables:

- `BUCKET_NAME` - the name of the Google Cloud Storage bucket to use (will throw if not set)
- `PORT_NUMBER` - the port to listen on (default: `1337`)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

Once spun, the webserver can be hit at the following HTTP endpoints:

- `/get`
  - `Content-Type`: `application/json`
  - `Body`: `{"key": "some-key"}`
- `/set`
  - `Content-Type`: `application/json`
  - `Body`: `{"key": "some-key", "value": "some-value"}`
- `/list`
  - `Content-Type`: `application/json`
  - `Body`: `{"prefix": "some-prefix", "flat": true}` (`flat` is optional, and defaults to `false`)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/amazing-feature`)
3. Commit your Changes (`git commit -m 'Add some amazing feature'`)
4. Push to the Branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the GPL v3.0 License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Author - [@0xCAP](https://twitter.com/0xCAP)

Project Link: [https://github.com/hood/coso](https://github.com/hood/coso)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->

[contributors-shield]: https://img.shields.io/github/contributors/hood/coso.svg?style=flat
[contributors-url]: https://github.com/hood/coso/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/hood/coso.svg?style=flat
[forks-url]: https://github.com/hood/coso/network/members
[stars-shield]: https://img.shields.io/github/stars/hood/coso.svg?style=flat
[stars-url]: https://github.com/hood/coso/stargazers
[issues-shield]: https://img.shields.io/github/issues/hood/coso.svg?style=flat
[issues-url]: https://github.com/hood/coso/issues
[license-shield]: https://img.shields.io/github/license/hood/coso.svg?style=flat
[license-url]: https://github.com/hood/coso/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png
