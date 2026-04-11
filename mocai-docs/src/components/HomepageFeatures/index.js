import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

const FeatureList = [
  {
    title: 'Easy To Start',
    Svg: require('@site/static/img/user_friendly_design.svg').default,
    description: (
      <>
        Install Mocaí with <code>go get</code>, create a <code>Mocker</code>, and start generating realistic test data with a small and practical API.
      </>
    ),
  },
  {
    title: 'Designed For Go Workflows',
    Svg: require('@site/static/img/streamlined_efficiency.svg').default,
    description: (
      <>
        Use Mocaí in Go projects with dependency injection through <code>MockGenerator</code>, functional options, and deterministic generation when your tests need stable output.
      </>
    ),
  },
  {
    title: 'Less Manual Fixture Work',
    Svg: require('@site/static/img/developed_in_go.svg').default,
    description: (
      <>
        Generate people, addresses, companies, CPF, certificates, national IDs, voter registration, and phone numbers without handcrafting repetitive mock data.
      </>
    ),
  },
];

function Feature({Svg, title, description}) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
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
