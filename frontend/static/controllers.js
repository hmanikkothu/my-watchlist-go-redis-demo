/**
 * Copyright 2016 The Kubernetes Authors All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

var redisApp = angular.module('redis', ['ui.bootstrap']);

/**
 * Constructor
 */
function RedisController() {}

RedisController.prototype.onRedis = function() {
    var val = this.scope_.item;
    if (!val) {
        return
    }
    var key = "k" + this.scope_.items.length
    var i = {ID: key, Value: val};
    this.scope_.items.push(i);
    this.scope_.item = "";
    this.http_.put("/item", i)
            .success(angular.bind(this, function(data) {
                this.scope_.redisResponse = "Updated.";
            }));
};

RedisController.prototype.onDelete = function(id, index) {
    console.log(id);

    this.scope_.items.splice(index, 1);
    this.http_.delete("/item?id=" + id)
    .success(angular.bind(this, function(data) {
        this.scope_.redisResponse = "Deleted.";
    }));
}

redisApp.controller('RedisCtrl', function ($scope, $http, $location) {
        $scope.controller = new RedisController();
        $scope.controller.scope_ = $scope;
        $scope.controller.location_ = $location;
        $scope.controller.http_ = $http;

        $scope.controller.http_.get("/allitems")
            .success(function(data) {
                console.log(data);
                $scope.items = data;
            });
});