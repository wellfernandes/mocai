import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

const FeatureList = [
  {
    title: 'User-Friendly Design',
    Svg: require('@site/static/img/user_friendly_design.svg').default,
    description: (
      <>
        Mocaí is designed to be easy to install and use, allowing you to set up your environment quickly and efficiently.
      </>
    ),
  },
  {
    title: 'Streamlined Efficiency',
    Svg: require('@site/static/img/streamlined_efficiency.svg').default,
    description: (
      <>
        With Mocaí, you can focus on developing your project while we handle the repetitive tasks. Streamline your workflow and enhance productivity.
      </>
    ),
  },
  {
    title: 'Developed in Go',
    Svg: require('@site/static/img/developed_in_go.svg').default,
    description: (
      <>
        Mocaí is built using Go, a powerful and efficient programming language. Leverage the performance and reliability of Go for your projects.
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
