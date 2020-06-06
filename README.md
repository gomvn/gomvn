GoMVN
=====

A lightweight self-hosted repository manager for your private Maven artifacts.


Installation
------------

Use docker to install this tool. Image is available at [Docker HUB](https://hub.docker.com/r/gomvn/gomvn).

For better accesibility, map these docker volumes:

| Path              | Description |
| ----------------- | ----------- |
| `/app/data`       | app data for persistency |
| `/app/config.yml` | configuration from outside of container, copy [default config](https://raw.githubusercontent.com/gomvn/gomvn/master/config.yml) |


User Guide
----------

On first run, admin account and his token is generated and prited into console.

You will need this to access [management api](https://gomvn.docs.apiary.io/), which is used to set user access.

If you don't have more users, you can use already created admin accout to deploy and access your maven artifacts.


### How to create java library 

Ensure that your `build.gradle` file contains configuration by this example:

```gradle
plugins {
    id 'maven-publish'
    id 'java'
}

publishing {
    repositories {
        maven {
            def releasesRepoUrl = "http://my-private-repository.example.com/release"
            def snapshotsRepoUrl = "http://my-private-repository.example.com/snapshot"
            name = 'mlj'
            url = project.version.endsWith('RELEASE') ? releasesRepoUrl : snapshotsRepoUrl
            credentials {
                username 'PUT HERE USERNAME'
                password 'PUT HERE TOKEN'
            }
        }
    }
    publications {
        maven(MavenPublication) {
            groupId = 'com.example'
            artifactId = 'library'
            version = '1.0.0.RELEASE'

            from components.java
        }
    }
}
```

### How to use your private maven repository

Append to your `build.gradle`:

```gradle
repositories {
    mavenCentral()
    maven {
        url "http://my-private-repository.example.com/release"
        credentials {
            username project.mljMavenUsername
            password project.mljMavenPassword
        }
    }
}

dependencies {
    implementation "com.example:library:1.0.0.RELEASE"
}
```
