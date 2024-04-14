import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';
import { useColorMode } from '@docusaurus/theme-common';
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import CodeBlock from "@theme/CodeBlock";
import Link from '@docusaurus/Link';
import useBaseUrl from '@docusaurus/useBaseUrl';
import About from './_about.mdx';
import HowItWorks from '../../../docs/how-it-works.mdx';

export default function HomepageFeatures(): JSX.Element {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          <div className="col col--6">
            <About></About>
          </div>
          <div className="col col--6">
            <HowItWorks></HowItWorks>
          </div>
        </div>
      </div>
    </section>
  );
}
