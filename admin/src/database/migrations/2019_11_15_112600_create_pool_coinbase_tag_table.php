<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreatePoolCoinbaseTagTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('pool_coinbase_tag', function (Blueprint $table) {
            $table->increments('id');
            $table->integer('pool_id')->index()->comment('矿池id');
            $table->string('tag')->unique()->comment('标记');
            $table->integer('created_at_ts')->index()->comment('创建时间戳');
            $table->integer('updated_at_ts')->index()->comment('更新时间戳');
            $table->timestamps();
            $table->unique(["pool_id", "tag"]);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('pool_coinbase_tag');
    }
}
