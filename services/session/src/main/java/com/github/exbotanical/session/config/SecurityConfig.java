package com.github.exbotanical.session.config;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.context.annotation.Bean;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.web.AuthenticationEntryPoint;
import org.springframework.security.web.SecurityFilterChain;
import org.springframework.security.web.access.AccessDeniedHandler;
import org.springframework.security.web.authentication.www.BasicAuthenticationFilter;
import org.springframework.security.web.header.HeaderWriter;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.exbotanical.session.config.filters.AuthenticationFilter;
import com.github.exbotanical.session.config.filters.CookieValidatorFilter;
import com.github.exbotanical.session.controllers.security.CookiesAuthenticationProvider;
import com.github.exbotanical.session.controllers.security.NotAuthorizedHandler;
import com.github.exbotanical.session.models.GouacheResponse;

@EnableWebSecurity
public class SecurityConfig {
  private static final ObjectMapper OBJECT_MAPPER = new ObjectMapper();

  private final String[] ALLOW_LIST = new String[] {
      "/user/login",
      "/user/register"
  };

  @Bean
  public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
    http
        .csrf().disable()

        .sessionManagement(session -> session
            .maximumSessions(1)
            .maxSessionsPreventsLogin(true))

        .addFilterBefore(new AuthenticationFilter(), BasicAuthenticationFilter.class)
        .addFilterBefore(new CookieValidatorFilter(), AuthenticationFilter.class)

        .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS)
        .and().logout().deleteCookies(CookieValidatorFilter.COOKIE_NAME)

        .and()

        .authorizeHttpRequests().antMatchers(ALLOW_LIST).permitAll()
        .anyRequest().authenticated()

        .and()
        .exceptionHandling()
        .authenticationEntryPoint(new AuthenticationEntryPoint() {
          @Override
          public void commence(HttpServletRequest request, HttpServletResponse response,
              AuthenticationException authException) throws IOException, ServletException {

            response.setStatus(HttpStatus.UNAUTHORIZED.value());
            response.setHeader(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE);

            GouacheResponse r =
                GouacheResponse.builder().friendly("You must be authenticated. @todo")
                    .internal(request.getRequestURI().toString()).build();

            System.out.println(r);

            OBJECT_MAPPER.writeValue(response.getOutputStream(), r);

          }
        })
        .accessDeniedHandler(accessDeniedHandler())

        .and()
        .headers()
        .addHeaderWriter(new HeaderWriter() {
          @Override
          public void writeHeaders(HttpServletRequest request, HttpServletResponse response) {
            response.setHeader("X-Powered-By", "gouache-session");
          }
        })

        .and()
        .httpBasic()

        .and()
        .authenticationProvider(new CookiesAuthenticationProvider())

    ;

    return http.build();
  }


  @Bean
  public AccessDeniedHandler accessDeniedHandler() {
    return new NotAuthorizedHandler();
  }

  @Bean
  public PasswordEncoder passwordEncoder() {
    return new BCryptPasswordEncoder(12);
  }
}
