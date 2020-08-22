<p align="center">
    <img alt="Agent Logo" src="/assets/logo.png" height="120" />
    <h3 align="center">Uptimedog Agent</h3>
    <p align="center">
        <a href="https://travis-ci.com/Uptimedog/Agent"><img src="https://travis-ci.com/Uptimedog/Agent.svg?branch=master"></a>
        <a href="https://github.com/Uptimedog/Agent/releases"><img src="https://img.shields.io/badge/Version-0.0.3-red.svg"></a>
        <a href="https://goreportcard.com/report/github.com/Uptimedog/Agent"><img src="https://goreportcard.com/badge/github.com/Uptimedog/Agent?v=0.0.1"></a>
        <a href="https://github.com/Uptimedog/Agent/blob/master/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg"></a>
    </p>
</p>


## Installation

Download [the latest agent binary.](https://github.com/Uptimedog/Agent/releases)

```zsh
$ curl -sL https://github.com/Uptimedog/Agent/releases/download/x.x.x/agent_x.x.x_OS.tar.gz | tar xz
```

Start the agent and it will auto register itself using your API key

```zsh
$ ./agent run --api_server uptimedog.io --api_key xxxxxxxx > agent.log
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Uptimedog Agent is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/uptimedog/agent/releases) for changelogs for each release version of the Agent. It contains summaries of the most noteworthy changes made in each release. Also see the [Milestones section](https://github.com/uptimedog/agent/milestones) for the future roadmap.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/uptimedog/agent/issues


## Security Issues

If you discover a security vulnerability within Uptimedog Agent, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2020, Uptimedog. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Uptimedog Agent** is authored and maintained by [@uptimedog](https://github.com/uptimedog).
