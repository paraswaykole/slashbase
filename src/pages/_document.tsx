import Document, { Html, Head, Main, NextScript } from 'next/document'

class SlashbaseDocument extends Document {

  render() {
    return (
      <Html lang="en">
        <Head>
        </Head>
        <body>
          <div className="appcontainer">
            <Main />
            <NextScript />
          </div>
        </body>
      </Html>
    )
  }

}

export default SlashbaseDocument
