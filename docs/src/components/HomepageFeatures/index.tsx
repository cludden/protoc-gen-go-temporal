import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';
import { useColorMode } from '@docusaurus/theme-common'

type FeatureItem = {
  title: string;
  Svg: React.ComponentType<React.ComponentProps<'svg'>>;
  description: JSX.Element;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'Define Workflow Schemas',
    Svg: require('@site/static/img/schema_black_24dp.svg').default,
    description: (
      <>
        Annotate your protobuf services and methods with Temporal options for Workflows,
        Activities, Queries, Signals, and Updates.
      </>
    ),
  },
  {
    title: 'Implement Worklows & Activities',
    Svg: require('@site/static/img/code_black_24dp.svg').default,
    description: (
      <>
        Generate Go code that includes types, methods, and functions for 
        building Temporal clients & workers and implement the required 
        Workflow and Activity interfaces
      </>
    ),
  },
  {
    title: 'Execute with Client & CLI',
    Svg: require('@site/static/img/rocket_launch_black_24dp.svg').default,
    description: (
      <>
        Run your Temporal worker using the generated helpers and interact with it using 
        the generated client, command line interface, or cross-namespace helpers
      </>
    ),
  },
];

function Feature({title, Svg, description}: FeatureItem) {
  // Custom code to change the fill color of the Cloud Arrow Up SVG
  // depending on if user is in dark/light mode
  let fill = 'black'
  const {colorMode} = useColorMode()
  if ( colorMode === 'dark') {
    fill = 'white' 
  }
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg fill={fill} className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures(): JSX.Element {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
