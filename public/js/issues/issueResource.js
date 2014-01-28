'use strict';

// Issue Resource
issues.factory("Issue", function ($resource) {
  return $resource(
    "/api/issues/:Id",
    {Id: "@id" },
    {
      "update": {
        method: "PUT",
        // transformRequest: function (data) {
          
        //   return angular.isObject(data) && String(data) !== '[object File]' ? param(data) : data;
        // }
      },
    }
  );
});