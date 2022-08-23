package com.github.exbotanical.session.entities;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.hibernate.validator.constraints.Length;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.validation.constraints.NotEmpty;

@Entity
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class Resource {
  @Id
  @GeneratedValue(strategy = GenerationType.AUTO)
  private Long id;

  // @todo verify max is inclusive
  @NotEmpty(message = "name is required")
  @Length(max = 32, message = "length must be less than or equal to 32 characters")
  private String name;

  @NotEmpty(message = "code is required")
  @Length(max = 64, message = "length must be less than or equal to 64 characters")
  private String code;

  private String data;
}
