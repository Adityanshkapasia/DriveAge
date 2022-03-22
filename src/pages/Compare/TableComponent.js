import React from "react";
import { connect } from "react-redux";
import { removeProduct } from "../actions/index";
import Col from "react-bootstrap/Col";
import Table from "react-bootstrap/Table";

class TableComponent extends React.Component {
  removeProduct = product => {
    this.props.removeProduct(product);
  };

  render() {
    return (
      <Col md={{ span: 10, offset: 1 }}>
        <Table hover className="tableProducts">
          <thead>
            <tr>
              <th>#</th>
              <th>Name</th>
              <th>Price(INR)</th>
              <th>Tank Capacity(l)</th>
              <th>Colour</th>
              <th>Fuel Economy</th>
            </tr>
          </thead>
          <tbody>
            {this.props.products.map(product => {
              return (
                <tr key={product.id}>
                  <td>{product.id}</td>
                  <td>{product.name}</td>
                  <td>{product.price}</td>
                  <td>{product.fuel}</td>
                  <td>{product.colour}</td>
                  <td>{product.capacity}</td>
                </tr>
              );
            })}
          </tbody>
        </Table>
      </Col>
    );
  }
}

const mapStateToProps = state => {
  return {
    products: state.products
  };
};

export default connect(
  mapStateToProps,
  { removeProduct }
)(TableComponent);
