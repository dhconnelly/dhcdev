# dhcdev

Source code and static content for my personal web site.

## Usage

    # run tests
    ./gradlew test

    # build the site and start the server
    ./gradlew run

    # now navigate to http://localhost:9000

## Run integration tests

    ./test.sh

## Deploy to production

The site is hosted on Digital Ocean. It will automatically pick up new pushes
to the main branch and deploy them after merge.

## License

All content in `pages` is licensed under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/).
Everything else falls under the MIT license (see the `LICENSE` file).
