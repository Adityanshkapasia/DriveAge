import React from "react";
import HeaderComponent from "./pages/Compare/HeaderComponent";
import ProductComponent from "./Compare/ProductComponent";
import TableComponent from "./Compare/TableComponent";
import "./Compare.css";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";

function App() {
  return (
    <div>
      <HeaderComponent />
      <Container>
        <br />
        <Row>
          <ProductComponent />
        </Row>
        <br />
        <Row>
          <TableComponent />
        </Row>
      </Container>
    </div>
  );
}

export default App;
