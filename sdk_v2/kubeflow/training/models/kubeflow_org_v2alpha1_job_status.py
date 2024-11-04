# coding: utf-8

"""
    Kubeflow Training OpenAPI Spec

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from kubeflow.training.configuration import Configuration


class KubeflowOrgV2alpha1JobStatus(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'active': 'int',
        'failed': 'int',
        'name': 'str',
        'ready': 'int',
        'succeeded': 'int',
        'suspended': 'int'
    }

    attribute_map = {
        'active': 'active',
        'failed': 'failed',
        'name': 'name',
        'ready': 'ready',
        'succeeded': 'succeeded',
        'suspended': 'suspended'
    }

    def __init__(self, active=0, failed=0, name='', ready=0, succeeded=0, suspended=0, local_vars_configuration=None):  # noqa: E501
        """KubeflowOrgV2alpha1JobStatus - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._active = None
        self._failed = None
        self._name = None
        self._ready = None
        self._succeeded = None
        self._suspended = None
        self.discriminator = None

        self.active = active
        self.failed = failed
        self.name = name
        self.ready = ready
        self.succeeded = succeeded
        self.suspended = suspended

    @property
    def active(self):
        """Gets the active of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501

        Active is the number of child Jobs with at least 1 pod in a running or pending state which are not marked for deletion.  # noqa: E501

        :return: The active of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501
        :rtype: int
        """
        return self._active

    @active.setter
    def active(self, active):
        """Sets the active of this KubeflowOrgV2alpha1JobStatus.

        Active is the number of child Jobs with at least 1 pod in a running or pending state which are not marked for deletion.  # noqa: E501

        :param active: The active of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501
        :type: int
        """
        if self.local_vars_configuration.client_side_validation and active is None:  # noqa: E501
            raise ValueError("Invalid value for `active`, must not be `None`")  # noqa: E501

        self._active = active

    @property
    def failed(self):
        """Gets the failed of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501

        Failed is the number of failed child Jobs.  # noqa: E501

        :return: The failed of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501
        :rtype: int
        """
        return self._failed

    @failed.setter
    def failed(self, failed):
        """Sets the failed of this KubeflowOrgV2alpha1JobStatus.

        Failed is the number of failed child Jobs.  # noqa: E501

        :param failed: The failed of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501
        :type: int
        """
        if self.local_vars_configuration.client_side_validation and failed is None:  # noqa: E501
            raise ValueError("Invalid value for `failed`, must not be `None`")  # noqa: E501

        self._failed = failed

    @property
    def name(self):
        """Gets the name of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501

        Name of the child Job.  # noqa: E501

        :return: The name of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this KubeflowOrgV2alpha1JobStatus.

        Name of the child Job.  # noqa: E501

        :param name: The name of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and name is None:  # noqa: E501
            raise ValueError("Invalid value for `name`, must not be `None`")  # noqa: E501

        self._name = name

    @property
    def ready(self):
        """Gets the ready of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501

        Ready is the number of child Jobs where the number of ready pods and completed pods is greater than or equal to the total expected pod count for the child Job.  # noqa: E501

        :return: The ready of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501
        :rtype: int
        """
        return self._ready

    @ready.setter
    def ready(self, ready):
        """Sets the ready of this KubeflowOrgV2alpha1JobStatus.

        Ready is the number of child Jobs where the number of ready pods and completed pods is greater than or equal to the total expected pod count for the child Job.  # noqa: E501

        :param ready: The ready of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501
        :type: int
        """
        if self.local_vars_configuration.client_side_validation and ready is None:  # noqa: E501
            raise ValueError("Invalid value for `ready`, must not be `None`")  # noqa: E501

        self._ready = ready

    @property
    def succeeded(self):
        """Gets the succeeded of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501

        Succeeded is the number of successfully completed child Jobs.  # noqa: E501

        :return: The succeeded of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501
        :rtype: int
        """
        return self._succeeded

    @succeeded.setter
    def succeeded(self, succeeded):
        """Sets the succeeded of this KubeflowOrgV2alpha1JobStatus.

        Succeeded is the number of successfully completed child Jobs.  # noqa: E501

        :param succeeded: The succeeded of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501
        :type: int
        """
        if self.local_vars_configuration.client_side_validation and succeeded is None:  # noqa: E501
            raise ValueError("Invalid value for `succeeded`, must not be `None`")  # noqa: E501

        self._succeeded = succeeded

    @property
    def suspended(self):
        """Gets the suspended of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501

        Suspended is the number of child Jobs which are in a suspended state.  # noqa: E501

        :return: The suspended of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501
        :rtype: int
        """
        return self._suspended

    @suspended.setter
    def suspended(self, suspended):
        """Sets the suspended of this KubeflowOrgV2alpha1JobStatus.

        Suspended is the number of child Jobs which are in a suspended state.  # noqa: E501

        :param suspended: The suspended of this KubeflowOrgV2alpha1JobStatus.  # noqa: E501
        :type: int
        """
        if self.local_vars_configuration.client_side_validation and suspended is None:  # noqa: E501
            raise ValueError("Invalid value for `suspended`, must not be `None`")  # noqa: E501

        self._suspended = suspended

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, KubeflowOrgV2alpha1JobStatus):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, KubeflowOrgV2alpha1JobStatus):
            return True

        return self.to_dict() != other.to_dict()