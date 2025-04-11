import '../styles/globals.css'
import * as React from 'react';
import Head from 'next/head';
import Page from '../components/Page';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';
import 'aos/dist/aos.css';
import {AppProps} from "next/app";
export default function App({ Component, pageProps }: AppProps) {
  return <React.Fragment>
    <Head>
      <meta
          name="viewport"
          content="width=device-width, initial-scale=1, shrink-to-fit=no"
      />
      <title>Coding Tasker</title>
    </Head>
    <Page>
      <Component {...pageProps} />
    </Page>
  </React.Fragment>
}
