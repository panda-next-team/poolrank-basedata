<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreatePoolTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('pool', function (Blueprint $table) {
            $table->increments('id');
            $table->string('name', "30")->index()->comment('名称');
            $table->string('website_url', "100")->comment('官网url');
            $table->string('icon')->comment('icon');
            $table->tinyInteger('status')->index()->default('1')->comment('状态');
            $table->integer('list_order')->index()->comment('自定义排序');
            $table->integer('created_at_ts')->index()->comment('创建时间戳');
            $table->integer('updated_at_ts')->index()->comment('更新时间戳');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('pool');
    }
}
