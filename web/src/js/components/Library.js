import React, { Component } from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'

class Library extends Component {
  render() {
    const { library, onClick } = this.props

    return (
      <div className="library">
        <List
          title={ ["title", "info"] }
          rows={ library }>
          <AddRow />
        </List>
      </div>
    );
  }
}
Library.propTypes = {
  library: PropTypes.array,
  onClick: PropTypes.func // todo
}

const mapStateToProps = (state) => ({
  library: [...state.rootReducer.library]
})

const mapDispatchToProps = {}

export default connect(mapStateToProps, mapDispatchToProps)(Library)


class List extends Component {
  render() {
    const { rows, title } = this.props

    return (
      <div className="list">
        <ListRow columns={title} />
        {rows.map((row, i) => (
          <ListRow key={i} columns={row} />
        ))}
        {this.props.children}
      </div>
    )
  }
}
List.propTypes = {
  rows: PropTypes.array
}

const ListRow = ({ columns }) => (
  <div className="row">
    {columns.map((column, i) => (
      <ListRowColumn key={i} item={column} />
    ))}
  </div>
)
ListRow.propTypes = {
  columns: PropTypes.array
}

const ListRowColumn = ({ item }) => (
    <div className="column">
      {item}
    </div>
)
ListRowColumn.propTypes = {
  // text: PropTypes.string
}

const AddRow = () => (
  <ListRow columns={[
    <input placeholder="title"></input>,
    <input placeholder="notes"></input>
  ]} />
)
