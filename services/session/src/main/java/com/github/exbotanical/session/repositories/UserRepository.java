package com.github.exbotanical.session.repositories;

import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBMapper;
import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBQueryExpression;
import com.amazonaws.services.dynamodbv2.model.AttributeValue;
import com.github.exbotanical.session.entities.User;
import com.github.exbotanical.session.errors.UserInvariantViolationException;
import com.github.exbotanical.session.errors.UserNotFoundException;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Repository;

@Repository
public class UserRepository {
  @Autowired
  private DynamoDBMapper dynamoDBMapper;

  @Value("${app.locale}")
  private String locale;

  private final Logger LOGGER = LoggerFactory.getLogger(UserRepository.class);


  public User getUserByUsername(String username) throws UserNotFoundException,
      UserInvariantViolationException {
    Map<String, AttributeValue> eav = new HashMap<>();
    eav.put(":v1", new AttributeValue().withS(username));

    DynamoDBQueryExpression<User> queryExpression = new DynamoDBQueryExpression<User>()
        .withKeyConditionExpression("Username = :v2")
        .withExpressionAttributeValues(eav);

    List<User> result = dynamoDBMapper.query(User.class, queryExpression);

    if (result.size() == 0) {
      throw new UserNotFoundException(username);
    } else if (result.size() > 1) {
      throw new UserInvariantViolationException(username);
    }

    return result.get(0);
  }

  public User createUser(User user) {
    dynamoDBMapper.save(user);
    // @todo use AOP
    LOGGER.info("saved user; value is now {}", user);

    return user;
  }

}
