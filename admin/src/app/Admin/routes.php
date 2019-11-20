<?php

use Illuminate\Routing\Router;

Admin::routes();

Route::group([
    'prefix'        => config('admin.route.prefix'),
    'namespace'     => config('admin.route.namespace'),
    'middleware'    => config('admin.route.middleware'),
], function (Router $router) {

    $router->resource('algorithms', AlgorithmController::class);
    $router->resource('pow_coins', PowCoinController::class);
    $router->resource('pools', PoolController::class);
    $router->get('/', 'HomeController@index')->name('admin.home');

});
