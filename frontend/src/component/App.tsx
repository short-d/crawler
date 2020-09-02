import React, {Component} from 'react';
import styles from './App.module.scss';
import {SampleAppGraphqlService} from '../service/sampleapp.graphql.service';
import {Change} from '../entity/change';
import {EnvService} from '../service/env.service';

interface IProps {
  sampleAppGraphQLService: SampleAppGraphqlService;
  envService: EnvService;
}

interface IState {
  env: string;
  changes: Change[];
}

export class App extends Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      env: '',
      changes: [],
    };
  }

  async componentDidMount() {
    const changeLog = await this.props.sampleAppGraphQLService.getChangeLog();
    const env = this.props.envService.getEnv().env;
    this.setState({
      changes: changeLog.changes,
      env
    });
  }

  render() {
    return (
      <div className={styles.App}>
        {this.state.env &&
        <div className={styles.Env}>
          {this.state.env}
        </div>
        }
        <div className={styles.Content}>
          <h1>Frontend Works!</h1>
          <h2>Change Log</h2>
          <ul className={styles.ChangeLog}>
            {this.state.changes.map(change =>
              <li>
                {`${change.title}`}
              </li>
            )}
          </ul>
        </div>
      </div>
    );
  }
}

export default App;
