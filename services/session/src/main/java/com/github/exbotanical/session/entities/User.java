package com.github.exbotanical.session.entities;

import java.sql.Date;

import javax.validation.constraints.NotEmpty;

import org.hibernate.validator.constraints.Length;

import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBAttribute;
import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBAutoGenerateStrategy;
import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBAutoGeneratedTimestamp;
import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBAutoGeneratedKey;
import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBHashKey;
import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBTable;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
@DynamoDBTable(tableName = "user")
public class User {

  @DynamoDBHashKey(attributeName = "Id")
  @DynamoDBAutoGeneratedKey
  private String id;

  // @todo composite key
  @NotEmpty(message = "username is required")
  @Length(max = 32, message = "length must be less than or equal to 32 characters")
  @DynamoDBAttribute(attributeName = "Username")
  private String username;

  @NotEmpty(message = "password is required")
  @Length(min = 6, message = "length must be greater than or equal to 6 characters")
  @Length(max = 32, message = "length must be less than or equal to 32 characters")
  @DynamoDBAttribute(attributeName = "PasswordHash")
  private String passwordHash;

  @DynamoDBAutoGeneratedTimestamp(strategy = DynamoDBAutoGenerateStrategy.CREATE)
  @DynamoDBAttribute(attributeName = "CreatedAt")
  private String createdAt;

  @DynamoDBAutoGeneratedTimestamp(strategy = DynamoDBAutoGenerateStrategy.ALWAYS)
  @DynamoDBAttribute(attributeName = "UpdatedAt")
  private String updatedAt;
}