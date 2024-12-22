import React, { useEffect, useState } from 'react';
import { Card, Table, Typography, Space, Button, InputNumber, Form } from 'antd';
import { useParams } from 'react-router-dom';

const { Title } = Typography;

const data = [
  {
    key: '1',
    product_id: 'P001',
    name: 'Smartphone',
    original_price: 499.99,
    seckill_price: 399.99,
    stock: 50,
  },
  {
    key: '2',
    product_id: 'P002',
    name: 'Laptop',
    original_price: 999.99,
    seckill_price: 799.99,
    stock: 30,
  },
];

function SeckillProductsPage() {
  const { eventId } = useParams(); // 获取路由参数
  const [quantity, setQuantity] = useState(1); // 选择的数量

  useEffect(() => {
    // 模拟从后端获取秒杀产品
    console.log('Fetching seckill products for event:', eventId);
  }, [eventId]);

  // 处理 "Buy Now" 按钮点击事件
  const handleBuyNow = (record) => {
    console.log('Buy now clicked for product:', record, 'Quantity:', quantity);
    // 在这里添加后端逻辑，例如创建秒杀订单
  };

  // 将 columns 移动到组件内部
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
      title: 'Original Price',
      dataIndex: 'original_price',
      key: 'original_price',
      render: (text) => `$${text.toFixed(2)}`,
    },
    {
      title: 'Seckill Price',
      dataIndex: 'seckill_price',
      key: 'seckill_price',
      render: (text) => `$${text.toFixed(2)}`,
    },
    {
      title: 'Stock',
      dataIndex: 'stock',
      key: 'stock',
    },
    {
      title: 'Actions',
      key: 'action',
      render: (_, record) => (
        <Space size="middle">
          <Button type="primary" onClick={() => handleBuyNow(record)}>
            Buy Now
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <div className="page-container">
      <Title level={2}>Seckill Products for Event {eventId}</Title>
      <Card title="Seckill Product List" bordered={false}>
        <Table
          columns={columns}
          dataSource={data}
          pagination={{ pageSize: 5 }}
        />
      </Card>
      <Form.Item style={{ marginTop: 16 }}>
        <InputNumber
          min={1}
          max={10}
          defaultValue={1}
          onChange={(value) => setQuantity(value)}
          style={{ width: '100px' }}
        />
        <span style={{ marginLeft: 8 }}>Select quantity</span>
      </Form.Item>
    </div>
  );
}

export default SeckillProductsPage;