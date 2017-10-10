import React, {PropTypes} from 'react';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import * as sessionActions from '../actions/sessionActions';
import TextInput from './TextInput';

class LoginPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {credentials: {username: '', password: ''}}
    this.onChange = this.onChange.bind(this);
    this.onSave = this.onSave.bind(this);
  }

  onChange(event) {
    return this.setState({credentials: {...this.state.credentials,
                                        [event.target.name]: event.target.value}});
  }

  onSave(event) {
    event.preventDefault();
    this.props.actions.loginUser(this.state.credentials);
  }

  render() {
    return (
      <div className="login page-component">
        <form>
          <TextInput
            autoFocus={true}
            name="username"
            label="username"
            placeholder=""
            value={this.state.credentials.username}
            onChange={this.onChange}/>

          <TextInput
            name="password"
            label="password"
            type="password"
            placeholder=""
            value={this.state.credentials.password}
            onChange={this.onChange}/>

          <input
            type="submit"
            value="log in"
            onClick={this.onSave} />
        </form>
      </div>
  );
  }
}

function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators(sessionActions, dispatch)
  };
}
export default connect(null, mapDispatchToProps)(LoginPage);
