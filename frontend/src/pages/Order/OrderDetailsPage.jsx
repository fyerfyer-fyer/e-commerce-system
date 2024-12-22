import React from 'react';
import { Card, Table, Typography, Button, Modal } from 'antd';

const { Title } = Typography;

const columns = [
  {
    title: 'Product ID',
    dataIndex: 'product_id',
    key: 'product_id',
  },
  {
    title: 'Quantity',
    dataIndex: 'quantity',
    key: 'quantity',
  },
  {
    title: 'Price',
    dataIndex: 'price',
    key: 'price',
    render: (text) => `$${text.toFixed(2)}`,
  },
];

const data = [
  {
    key: '1',
    product_id: 'P001',
    quantity: 2,
    price: 19.99,
  },
  {
    key: '2',
    product_id: 'P002',
    quantity: 1,
    price: 49.99,
  },
];

function OrderDetailsPage() {
  const orderId = '12345'; // This should be dynamically fetched from the URL or state

  return (
    <div className="page-container">
      <Title level={2}>Order Details</Title>
      <Card title={`Order #${orderId}`} bordered={false}>
        <Table columns={columns} dataSource={data} pagination={false} />
        <div style={{ marginTop: 16 }}>
          <strong>Total Amount:</strong> $69.98
        </div>
        <div style={{ marginTop: 16 }}>
          <Button type="primary" danger onClick={() => showCancelConfirm()}>
            Cancel Order
          </Button>
        </div>
      </Card>
    </div>
  );

  function showCancelConfirm() {
    Modal.confirm({
      title: 'Cancel Order',
      content: `Are you sure you want to cancel order #${orderId}?`,
      okText: 'Cancel',
      okType: 'danger',
      cancelText: 'No',
      onOk() {
        console.log('Order canceled:', orderId);
        // Here you can add the logic to cancel the order on the backend
      },
    });
  }
}

export default OrderDetailsPage;