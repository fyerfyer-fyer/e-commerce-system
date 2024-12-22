import React from 'react';
import { Card, Table, Typography, Space, Button, Modal, message } from 'antd';
import { Link } from 'react-router-dom';

const { Title } = Typography;

const data = [
  {
    key: '1',
    product_id: 'P001',
    name: 'Smartphone',
    price: 499.99,
    stock: 100,
    category: 'electronics',
  },
  {
    key: '2',
    product_id: 'P002',
    name: 'Laptop',
    price: 999.99,
    stock: 50,
    category: 'electronics',
  },
];

function ProductListPage() {
  function showDeleteConfirm(record) {
    Modal.confirm({
      title: 'Delete Product',
      content: `Are you sure you want to delete the product "${record.name}"?`,
      okText: 'Delete',
      okType: 'danger',
      cancelText: 'Cancel',
      onOk() {
        console.log('Deleted product:', record);
        // Here you can add the logic to delete the product on the backend
        message.success('Product deleted successfully!');
      },
    });
  }

  // 将 columns 放到组件内部，这样可以直接访问 showDeleteConfirm 函数
  const columns = [
    {
      title: 'Product ID',
      dataIndex: 'product_id',
      key: 'product_id',
    },
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: 'Price',
      dataIndex: 'price',
      key: 'price',
      render: (text) => `$${text.toFixed(2)}`,
    },
    {
      title: 'Stock',
      dataIndex: 'stock',
      key: 'stock',
    },
    {
      title: 'Category',
      dataIndex: 'category',
      key: 'category',
    },
    {
      title: 'Actions',
      key: 'action',
      render: (_, record) => (
        <Space size="middle">
          <Link to={`/admin/product/edit/${record.product_id}`}>Edit</Link>
          <a href="#" onClick={() => showDeleteConfirm(record)}>Delete</a>
        </Space>
      ),
    },
  ];

  return (
    <div className="page-container">
      <Title level={2}>Product List</Title>
      <Card title="Product Management" bordered={false}>
        <Table columns={columns} dataSource={data} pagination={{ pageSize: 5 }} />
        <Button type="primary" style={{ marginTop: 16 }} onClick={() => window.location.href = '/admin/product/add'}>
          Add New Product
        </Button>
      </Card>
    </div>
  );
}

export default ProductListPage;