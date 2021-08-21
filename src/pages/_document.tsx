import Document, { Html, Head, Main, NextScript } from 'next/document'

class SlashbaseDocument extends Document {

  render() {
    return (
      <Html lang="en">
        <Head>
          <meta charSet="UTF-8" />
          <meta name='theme-color' content='#ff8f6d'/>
          <link rel="preconnect" href="https://fonts.gstatic.com"/>
          <link href="https://fonts.googleapis.com/css2?family=PT+Sans:ital,wght@0,400;0,700;1,400;1,700&family=Source+Sans+Pro:ital,wght@0,200;0,300;0,400;0,600;0,700;1,200;1,300;1,400;1,600;1,700&display=swap" rel="stylesheet" />
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
