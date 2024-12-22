import React, { useState } from 'react';
import { Card, Table, Button, Modal, Form, Input, Space } from 'antd';

const data = [
  {
    key: '1',
    address_line1: '123 Main St',
    city: 'New York',
    state: 'NY',
    postal_code: '10001',
    country: 'USA',
  },
  {
    key: '2',
    address_line1: '456 Elm St',
    city: 'Los Angeles',
    state: 'CA',
    postal_code: '90001',
    country: 'USA',
  },
];

function UserAddressPage() {
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [selectedAddress, setSelectedAddress] = useState(null);
  const [form] = Form.useForm();

  const showEditModal = (record) => {
    setIsModalVisible(true);
    setSelectedAddress(record);
    form.setFieldsValue(record);
  };

  const handleOk = () => {
    form.validateFields().then((values) => {
      console.log('Updated address:', values);
      // Here you can add the logic to update or add the address on the backend
      setIsModalVisible(false);
      setSelectedAddress(null);
      form.resetFields();
    });
  };

  const handleCancel = () => {
    setIsModalVisible(false);
    setSelectedAddress(null);
    form.resetFields();
  };

  const showDeleteConfirm = (record) => {
    Modal.confirm({
      title: 'Delete Address',
      content: `Are you sure you want to delete the address "${record.address_line1}"?`,
      okText: 'Delete',
      okType: 'danger',
      cancelText: 'Cancel',
      onOk() {
        console.log('Deleted address:', record);
        // Here you can add the logic to delete the address on the backend
      },
    });
  };

  const columns = [
    {
      title: 'Address Line 1',
      dataIndex: 'address_line1',
      key: 'address_line1',
    },
    {
      title: 'City',
      dataIndex: 'city',
      key: 'city',
    },
    {
      title: 'State/Province',
      dataIndex: 'state',
      key: 'state',
    },
    {
      title: 'Postal Code',
      dataIndex: 'postal_code',
      key: 'postal_code',
    },
    {
      title: 'Country',
      dataIndex: 'country',
      key: 'country',
    },
    {
      title: 'Actions',
      key: 'action',
      render: (_, record) => (
        <Space size="middle">
          <Button type="link" onClick={() => showEditModal(record)}>
            Edit
          </Button>
          <Button type="link" danger onClick={() => showDeleteConfirm(record)}>
            Delete
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <div className="page-container">
      <Card title="User Addresses" bordered={false}>
        <Table columns={columns} dataSource={data} pagination={{ pageSize: 5 }} />
        <Button type="primary" style={{ marginTop: 16 }} onClick={() => showEditModal({})}>
          Add New Address
        </Button>
        <Modal
          title={selectedAddress ? 'Edit Address' : 'Add New Address'}
          visible={isModalVisible}
          onOk={handleOk}
          onCancel={handleCancel}
        >
          <Form form={form} layout="vertical">
            <Form.Item
              name="address_line1"
              label="Address Line 1"
              rules={[{ required: true, message: 'Please enter the address line 1!' }]}
            >
              <Input placeholder="Enter address line 1" />
            </Form.Item>
            <Form.Item
              name="address_line2"
              label="Address Line 2"
            >
              <Input placeholder="Enter address line 2 (optional)" />
            </Form.Item>
            <Form.Item
              name="city"
              label="City"
              rules={[{ required: true, message: 'Please enter the city!' }]}
            >
              <Input placeholder="Enter city" />
            </Form.Item>
            <Form.Item
              name="state"
              label="State/Province"
              rules={[{ required: true, message: 'Please enter the state/province!' }]}
            >
              <Input placeholder="Enter state/province" />
            </Form.Item>
            <Form.Item
              name="postal_code"
              label="Postal Code"
              rules={[{ required: true, message: 'Please enter the postal code!' }]}
            >
              <Input placeholder="Enter postal code" />
            </Form.Item>
            <Form.Item
              name="country"
              label="Country"
              rules={[{ required: true, message: 'Please enter the country!' }]}
            >
              <Input placeholder="Enter country" />
            </Form.Item>
            <Form.Item
              name="is_default"
              valuePropName="checked"
              initialValue={false}
            >
              <Input.Checkbox>Set as Default Address</Input.Checkbox>
            </Form.Item>
          </Form>
        </Modal>
      </Card>
    </div>
  );
}

export default UserAddressPage;