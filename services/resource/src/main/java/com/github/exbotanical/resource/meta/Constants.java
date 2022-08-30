package com.github.exbotanical.resource.meta;

public class Constants {
  public static final String E_RESOURCE_UPDATE_FMT = "An exception occurred while updating the resource with id %s";

  public static final String E_RESOURCE_DELETE_FMT = "An exception occurred while deleting the resource with id %s";

  public static final String E_USER_DUPE_INVARIANT_FMT = "User with username %s was duplicated";

  public static final String E_SESSION_EXPIRED_FMT = "Session with id %s for user %s expired on %s";

  public static final String E_SESSION_NOT_FOUND_FMT = "No session found for id %s";

  public static final String E_ROUTE_NOT_FOUND_FMT = "The requested route %s does not exist.";

  public static final String E_METHOD_NOT_ALLOWED_FMT = "The requested method %s at route %s is not supported.";

  public static final String E_INVALID_INPUT = "The provided input was not valid.";

  public static final String E_GENERIC = "An exception occurred and the requested operation could not be completed.";

  public static final String E_UNAUTHORIZED = "You are not authorized to access this resource.";

  public static final String E_SESSION_ID_NOT_FOUND = "No session Id found";

  public static final String E_COOKIE_NOT_FOUND = "No cookie found";

  public static final String E_INTERNAL_SERVER_ERROR = "An internal server exception has occurred.";
}
