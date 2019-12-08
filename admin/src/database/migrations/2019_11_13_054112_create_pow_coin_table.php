<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreatePowCoinTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('pow_coin', function (Blueprint $table) {
            $table->increments('id');
            $table->string('name', "15")->unique()->comment('中文名称');
            $table->string('en_name', "15")->unique()->comment('英文名称');
            $table->string('en_tag', "10")->unique()->comment('英文缩写');
            $table->double('max_supply')->comment('最大供应总量');
            $table->integer('algorithm_id')->index()->comment('算法id');
            $table->date('release_date')->comment('发行日期');
            $table->integer('block_time')->comment('理论出块时间(秒)');
            $table->string('website_url')->comment('官网 url');
            $table->string('github_url')->comment('github url');
            $table->string('icon')->comment('icon');
            $table->string('intro')->comment('简介');
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
        Schema::dropIfExists('pow_coin');
    }
}
