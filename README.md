<div id="top"></div>
<!--
*** Thanks for checking out the consumer_rmq_fsevent. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
<h3 align="center">consumer_rmq_fsevent</h3>

  <p align="center">
    project_description
    <br />
    <br />
    <br />
    ·
    <a href="https://github.com/surendratiwari3/consumer_rmq_fsevent/issues">Report Bug</a>
    ·
    <a href="https://github.com/surendratiwari3/consumer_rmq_fsevent/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
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
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

consumer_rmq_fsevent: It will consume the freeswitch conference event from rabbitmq and generate the twilio like statuscallback event. You can also send the genrated statuscallback event to configured statuscallback url. Please set the confvar variable for making consumer_rmq_fsevent to send statuscallback to http url. 

<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [FreeSWITCH](https://freeswitch.com/)
* [GoLang](https://go.dev/)
* [RabbitMq](https://www.rabbitmq.com/)
* [Redis](https://redis.io/)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started
consumer_rmq_fsevent is rabbitmq consumer written in golang. It will consume the freeswitch conference events, translate them to twilio statuscallback events and send to configured statuscallback url .
You can configure the below channel variables before create conference :
confvar_statuscallback_url
confvar_statuscallback_method


### Prerequisites

* rabbitmq for freeswitch events
* redis for sequence for each event
* golang for build and run the consumer_rmq_fsevent

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/surendratiwari3/consumer_rmq_fsevent.git
   ```
3. Change directory to clone repository and Install go dependency
   ```sh
   cd consumer_rmq_fsevent
   go mod download
   ```
4. Compile and build the package 
   ```sh
   go build -o consumer_rmq_fsevent main.go
   ```
5. Run the package
   ```sh
   ./consumer_rmq_fsevent
   ```

<p align="right">(<a href="#top">back to top</a>)</p>


## Roadmap

- [] Announcement Event Support
- [] Participant Hold/Unhold/Mute/UnMute
- [] Conference Recording Event Support
- [] Dial Events support

See the [open issues](https://github.com/surendratiwari3/consumer_rmq_fsevent/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Project Link: [https://github.com/surendratiwari3/consumer_rmq_fsevent](https://github.com/surendratiwari3/consumer_rmq_fsevent)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* []()
* []()
* []()

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/surendratiwari3/consumer_rmq_fsevent.svg?style=for-the-badge
[contributors-url]: https://github.com/surendratiwari3/consumer_rmq_fsevent/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/surendratiwari3/consumer_rmq_fsevent.svg?style=for-the-badge
[forks-url]: https://github.com/surendratiwari3/consumer_rmq_fsevent/network/members
[stars-shield]: https://img.shields.io/github/stars/surendratiwari3/consumer_rmq_fsevent.svg?style=for-the-badge
[stars-url]: https://github.com/surendratiwari3/consumer_rmq_fsevent/stargazers
[issues-shield]: https://img.shields.io/github/issues/surendratiwari3/consumer_rmq_fsevent.svg?style=for-the-badge
[issues-url]: https://github.com/surendratiwari3/consumer_rmq_fsevent/issues
[license-shield]: https://img.shields.io/github/license/surendratiwari3/consumer_rmq_fsevent.svg?style=for-the-badge
[license-url]: https://github.com/surendratiwari3/consumer_rmq_fsevent/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/surendra-tiwari-st-ab569a15

