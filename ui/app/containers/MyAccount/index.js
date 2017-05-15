
import React from 'react';
import { connect } from 'react-redux';
import { createStructuredSelector } from 'reselect';
import { makeSelectUser } from '../App/selectors';

export class MyAccount extends React.PureComponent {
  render() {
    return (
      <h1>
            Hello {this.props.user.name}
      </h1>
    );
  }
}

export default connect(createStructuredSelector({
    user:  makeSelectUser(),
}))(MyAccount);
